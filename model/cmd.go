package model

import (
	"github.com/GlennLiu0607/glennctl/internal/cobrax"
	"github.com/GlennLiu0607/glennctl/model/mongo"
	"github.com/GlennLiu0607/glennctl/model/sql/command"
)

var (
    // Cmd describes a model command.
    Cmd             = cobrax.NewCommand("model")
    mysqlCmd        = cobrax.NewCommand("mysql")
    ddlCmd          = cobrax.NewCommand("ddl", cobrax.WithRunE(command.MysqlDDL))
    datasourceCmd   = cobrax.NewCommand("datasource", cobrax.WithRunE(command.MySqlDataSource))
    pgCmd           = cobrax.NewCommand("pg", cobrax.WithRunE(command.PostgreSqlDataSource))
    pgDatasourceCmd = cobrax.NewCommand("datasource", cobrax.WithRunE(command.PostgreSqlDataSource))
    dmCmd           = cobrax.NewCommand("dm", cobrax.WithRunE(command.DamengDataSource))
    dmDatasourceCmd = cobrax.NewCommand("datasource", cobrax.WithRunE(command.DamengDataSource))
    mongoCmd        = cobrax.NewCommand("mongo", cobrax.WithRunE(mongo.Action))
)

func init() {
    var (
        ddlCmdFlags          = ddlCmd.Flags()
        datasourceCmdFlags   = datasourceCmd.Flags()
        pgDatasourceCmdFlags = pgDatasourceCmd.Flags()
        dmDatasourceCmdFlags = dmDatasourceCmd.Flags()
        mongoCmdFlags        = mongoCmd.Flags()
    )

	ddlCmdFlags.StringVarP(&command.VarStringSrc, "src", "s")
	ddlCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
	ddlCmdFlags.StringVar(&command.VarStringStyle, "style")
	ddlCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
	ddlCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
	ddlCmdFlags.StringVar(&command.VarStringDatabase, "database")
	ddlCmdFlags.StringVar(&command.VarStringHome, "home")
	ddlCmdFlags.StringVar(&command.VarStringRemote, "remote")
	ddlCmdFlags.StringVar(&command.VarStringBranch, "branch")

	datasourceCmdFlags.StringVar(&command.VarStringURL, "url")
	datasourceCmdFlags.StringSliceVarP(&command.VarStringSliceTable, "table", "t")
	datasourceCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
	datasourceCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
	datasourceCmdFlags.StringVar(&command.VarStringStyle, "style")
	datasourceCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
	datasourceCmdFlags.StringVar(&command.VarStringHome, "home")
	datasourceCmdFlags.StringVar(&command.VarStringRemote, "remote")
	datasourceCmdFlags.StringVar(&command.VarStringBranch, "branch")

    pgDatasourceCmdFlags.StringVar(&command.VarStringURL, "url")
    pgDatasourceCmdFlags.StringSliceVarP(&command.VarStringSliceTable, "table", "t")
    pgDatasourceCmdFlags.StringVarPWithDefaultValue(&command.VarStringSchema, "schema", "s", "public")
    pgDatasourceCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
    pgDatasourceCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
    pgDatasourceCmdFlags.StringVar(&command.VarStringStyle, "style")
    pgDatasourceCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
    pgDatasourceCmdFlags.BoolVar(&command.VarBoolStrict, "strict")
    pgDatasourceCmdFlags.StringVar(&command.VarStringHome, "home")
    pgDatasourceCmdFlags.StringVar(&command.VarStringRemote, "remote")
    pgDatasourceCmdFlags.StringVar(&command.VarStringBranch, "branch")
    pgCmd.PersistentFlags().StringSliceVarPWithDefaultValue(&command.VarStringSliceIgnoreColumns,
        "ignore-columns", "i", []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"})
    pgCmd.PersistentFlags().StringVarPWithDefaultValue(&command.VarStringCachePrefix, "prefix", "p", "cache")

    dmDatasourceCmdFlags.StringVar(&command.VarStringURL, "url")
    dmDatasourceCmdFlags.StringSliceVarP(&command.VarStringSliceTable, "table", "t")
    dmDatasourceCmdFlags.StringVarPWithDefaultValue(&command.VarStringSchema, "schema", "s", "SYSDBA")
    dmDatasourceCmdFlags.BoolVarP(&command.VarBoolCache, "cache", "c")
    dmDatasourceCmdFlags.StringVarP(&command.VarStringDir, "dir", "d")
    dmDatasourceCmdFlags.StringVar(&command.VarStringStyle, "style")
    dmDatasourceCmdFlags.BoolVar(&command.VarBoolIdea, "idea")
    dmDatasourceCmdFlags.BoolVar(&command.VarBoolStrict, "strict")
    dmDatasourceCmdFlags.StringVar(&command.VarStringHome, "home")
    dmDatasourceCmdFlags.StringVar(&command.VarStringRemote, "remote")
    dmDatasourceCmdFlags.StringVar(&command.VarStringBranch, "branch")
    dmCmd.PersistentFlags().StringSliceVarPWithDefaultValue(&command.VarStringSliceIgnoreColumns,
        "ignore-columns", "i", []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"})
    dmCmd.PersistentFlags().StringVarPWithDefaultValue(&command.VarStringCachePrefix, "prefix", "p", "cache")

	mongoCmdFlags.StringSliceVarP(&mongo.VarStringSliceType, "type", "t")
	mongoCmdFlags.BoolVarP(&mongo.VarBoolCache, "cache", "c")
	mongoCmdFlags.StringVarP(&mongo.VarStringPrefix, "prefix", "p")
	mongoCmdFlags.BoolVarP(&mongo.VarBoolEasy, "easy", "e")
	mongoCmdFlags.StringVarP(&mongo.VarStringDir, "dir", "d")
	mongoCmdFlags.StringVar(&mongo.VarStringStyle, "style")
	mongoCmdFlags.StringVar(&mongo.VarStringHome, "home")
	mongoCmdFlags.StringVar(&mongo.VarStringRemote, "remote")
	mongoCmdFlags.StringVar(&mongo.VarStringBranch, "branch")

	mysqlCmd.PersistentFlags().BoolVar(&command.VarBoolStrict, "strict")
	mysqlCmd.PersistentFlags().StringSliceVarPWithDefaultValue(&command.VarStringSliceIgnoreColumns,
		"ignore-columns", "i", []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"})
	mysqlCmd.PersistentFlags().StringVarPWithDefaultValue(&command.VarStringCachePrefix, "prefix", "p", "cache")

    mysqlCmd.AddCommand(datasourceCmd, ddlCmd)
    pgCmd.AddCommand(pgDatasourceCmd)
    dmCmd.AddCommand(dmDatasourceCmd)
    Cmd.AddCommand(mysqlCmd, mongoCmd, pgCmd, dmCmd)
}
