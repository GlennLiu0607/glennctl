package util

import (
    "bytes"
    goformat "go/format"
    "os"
    "regexp"
    "time"
    "text/template"

    "github.com/glenn/glennctl/internal/errorx"
    "github.com/glenn/glennctl/util/pathx"
)

const regularPerm = 0o666

// DefaultTemplate is a tool to provides the text/template operations
type DefaultTemplate struct {
	name  string
	text  string
	goFmt bool
}

// With returns an instance of DefaultTemplate
func With(name string) *DefaultTemplate {
	return &DefaultTemplate{
		name: name,
	}
}

// Parse accepts a source template and returns DefaultTemplate
func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	t.text = text
	return t
}

// GoFmt sets the value to goFmt and marks the generated codes will be formatted or not
func (t *DefaultTemplate) GoFmt(format bool) *DefaultTemplate {
	t.goFmt = format
	return t
}

// SaveTo writes the codes to the target path
func (t *DefaultTemplate) SaveTo(data any, path string, forceUpdate bool) error {
    if pathx.FileExists(path) && !forceUpdate {
        return nil
    }

    // enrich data with common header variables when possible
    data = enrichHeaderData(data, path)

    output, err := t.Execute(data)
    if err != nil {
        return err
    }

    return os.WriteFile(path, output.Bytes(), regularPerm)
}

// Execute returns the codes after the template executed
func (t *DefaultTemplate) Execute(data any) (*bytes.Buffer, error) {
    // enrich data with common header variables (FilePath unknown here)
    data = enrichHeaderData(data, "")

    tem, err := template.New(t.name).Parse(t.text)
    if err != nil {
        return nil, errorx.Wrap(err, "template parse error:", t.text)
    }

	buf := new(bytes.Buffer)
	if err = tem.Execute(buf, data); err != nil {
		return nil, errorx.Wrap(err, "template execute error:", t.text)
	}

	if !t.goFmt {
		return buf, nil
	}

	formatOutput, err := goformat.Source(buf.Bytes())
	if err != nil {
		return nil, errorx.Wrap(err, "go format error:", buf.String())
	}

	buf.Reset()
	buf.Write(formatOutput)
	return buf, nil
}

// IsTemplateVariable returns true if the text is a template variable.
// The text must start with a dot and be a valid template.
func IsTemplateVariable(text string) bool {
	match, _ := regexp.MatchString(`(?m)^{{(\.\w+)+}}$`, text)
	return match
}

// TemplateVariable returns the variable name of the template.
func TemplateVariable(text string) string {
    if IsTemplateVariable(text) {
        return text[3 : len(text)-2]
    }
    return ""
}

// enrichHeaderData adds shared header variables into a map-based template data.
// Recognized keys: Date, LastEditors, LastEditTime, FilePath
func enrichHeaderData(data any, filePath string) any {
    now := time.Now().Format("2006-01-02 15:04:05")

    switch d := data.(type) {
    case map[string]any:
        if _, ok := d["Date"]; !ok {
            d["Date"] = now
        }
        if _, ok := d["LastEditTime"]; !ok {
            d["LastEditTime"] = now
        }
        if _, ok := d["LastEditors"]; !ok {
            d["LastEditors"] = "glennctl"
        }
        if filePath != "" {
            d["FilePath"] = filePath
        } else if _, ok := d["FilePath"]; !ok {
            d["FilePath"] = ""
        }
        return d
    case map[string]string:
        if _, ok := d["Date"]; !ok {
            d["Date"] = now
        }
        if _, ok := d["LastEditTime"]; !ok {
            d["LastEditTime"] = now
        }
        if _, ok := d["LastEditors"]; !ok {
            d["LastEditors"] = "glennctl"
        }
        if filePath != "" {
            d["FilePath"] = filePath
        } else if _, ok := d["FilePath"]; !ok {
            d["FilePath"] = ""
        }
        return d
    default:
        return data
    }
}
