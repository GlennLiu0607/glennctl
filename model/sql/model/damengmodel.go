package model

import (
    "database/sql"
    "strings"

    "gorm.io/gorm"
)

// DamengModel introspects Dameng DM8 schema using Oracle-compatible views
type DamengModel struct {
    db *gorm.DB
}

// NewDamengModel creates a new DamengModel
func NewDamengModel(db *gorm.DB) *DamengModel {
    return &DamengModel{db: db}
}

// GetAllTables returns all table names under a schema (owner)
func (m *DamengModel) GetAllTables(schema string) ([]string, error) {
    owner := strings.ToUpper(schema)
    var tables []string
    // Prefer ALL_TABLES to support schema filtering
    rows := m.db.Raw(`select table_name from all_tables where owner = ?`, owner)
    if rows.Error != nil {
        return nil, rows.Error
    }
    type rec struct{ TableName string `db:"TABLE_NAME"` }
    var list []rec
    if err := rows.Scan(&list).Error; err != nil {
        return nil, err
    }
    for _, r := range list {
        tables = append(tables, r.TableName)
    }
    return tables, nil
}

// FindColumns returns ColumnData for a given schema and table
func (m *DamengModel) FindColumns(schema, table string) (*ColumnData, error) {
    owner := strings.ToUpper(schema)
    tbl := strings.ToUpper(table)

    // Columns
    type colRec struct {
        ColumnName      string         `db:"COLUMN_NAME"`
        DataType        string         `db:"DATA_TYPE"`
        ColumnDefault   sql.NullString `db:"DATA_DEFAULT"`
        Nullable        string         `db:"NULLABLE"`
        OrdinalPosition int            `db:"COLUMN_ID"`
        Comment         sql.NullString `db:"COMMENTS"`
    }

    var cols []colRec
    columnsQuery := `select c.column_name, c.data_type, c.data_default, c.nullable, c.column_id,
       (select comments from all_col_comments cc where cc.owner = c.owner and cc.table_name = c.table_name and cc.column_name = c.column_name) as comments
     from all_tab_columns c where c.owner = ? and c.table_name = ? order by c.column_id`
    if err := m.db.Raw(columnsQuery, owner, tbl).Scan(&cols).Error; err != nil {
        return nil, err
    }

    // Indexes (including primary key)
    type idxRec struct {
        IndexName    string         `db:"INDEX_NAME"`
        Uniqueness   string         `db:"UNIQUENESS"`
        ColumnName   string         `db:"COLUMN_NAME"`
        ColumnPos    int            `db:"COLUMN_POSITION"`
        IsPrimaryKey sql.NullString `db:"IS_PRIMARY"`
    }

    var idxs []idxRec
    indexQuery := `select i.index_name, i.uniqueness, ic.column_name, ic.column_position,
       (case when exists (
            select 1 from all_constraints ac where ac.owner = i.table_owner and ac.table_name = i.table_name and ac.constraint_type = 'P' and ac.index_name = i.index_name
        ) then 'Y' else 'N' end) as is_primary
     from all_indexes i join all_ind_columns ic
       on i.index_name = ic.index_name and i.table_owner = ic.table_owner
     where i.table_owner = ? and i.table_name = ?
     order by ic.column_position`
    if err := m.db.Raw(indexQuery, owner, tbl).Scan(&idxs).Error; err != nil {
        return nil, err
    }

    // Build index map by column
    indexMap := make(map[string][]*DbIndex)
    var primaryCols []string
    for _, rec := range idxs {
        if rec.IsPrimaryKey.Valid && strings.ToUpper(rec.IsPrimaryKey.String) == "Y" {
            primaryCols = append(primaryCols, rec.ColumnName)
            indexMap[rec.ColumnName] = append(indexMap[rec.ColumnName], &DbIndex{
                IndexName:  indexPri,
                SeqInIndex: rec.ColumnPos,
            })
            continue
        }
        nonUnique := 1
        if strings.ToUpper(rec.Uniqueness) == "UNIQUE" {
            nonUnique = 0
        }
        indexMap[rec.ColumnName] = append(indexMap[rec.ColumnName], &DbIndex{
            IndexName:  rec.IndexName,
            NonUnique:  nonUnique,
            SeqInIndex: rec.ColumnPos,
        })
    }

    // Assemble Column list
    var list []*Column
    for _, c := range cols {
        isNullAble := "YES"
        if strings.ToUpper(c.Nullable) == "N" {
            isNullAble = "NO"
        }
        var dft any
        if c.ColumnDefault.Valid {
            dft = c.ColumnDefault
        }

        dt := convertDamengTypeIntoMysqlType(c.DataType)
        if len(indexMap[c.ColumnName]) > 0 {
            for _, i := range indexMap[c.ColumnName] {
                list = append(list, &Column{
                    DbColumn: &DbColumn{
                        Name:            c.ColumnName,
                        DataType:        dt,
                        Comment:         c.Comment.String,
                        ColumnDefault:   dft,
                        IsNullAble:      isNullAble,
                        OrdinalPosition: c.OrdinalPosition,
                    },
                    Index: i,
                })
            }
        } else {
            list = append(list, &Column{
                DbColumn: &DbColumn{
                    Name:            c.ColumnName,
                    DataType:        dt,
                    Comment:         c.Comment.String,
                    ColumnDefault:   dft,
                    IsNullAble:      isNullAble,
                    OrdinalPosition: c.OrdinalPosition,
                },
            })
        }
    }

    // Build ColumnData
    var columnData ColumnData
    columnData.Db = owner
    columnData.Table = tbl
    columnData.Columns = list
    return &columnData, nil
}

func convertDamengTypeIntoMysqlType(in string) string {
    tp := strings.ToLower(in)
    // Normalize Oracle-like types to MySQL-ish names used by converter
    switch {
    case strings.HasPrefix(tp, "varchar2"):
        return "varchar2"
    case strings.HasPrefix(tp, "nvarchar2"):
        return "nvarchar2"
    case strings.HasPrefix(tp, "number"):
        return "number"
    case strings.HasPrefix(tp, "timestamp"):
        return "timestamp"
    case tp == "date":
        return "date"
    case tp == "integer":
        return "int"
    case tp == "clob":
        return "text"
    case tp == "blob":
        return "blob"
    case tp == "nchar":
        return "nchar"
    case tp == "char":
        return "char"
    default:
        return tp
    }
}