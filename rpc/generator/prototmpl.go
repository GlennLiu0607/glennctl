package generator

import (
    _ "embed"
    "path/filepath"
    "strings"
    "time"

    "github.com/glenn/glennctl/util"
    "github.com/glenn/glennctl/util/pathx"
    "github.com/glenn/glennctl/util/stringx"
)

//go:embed rpc.tpl
var rpcTemplateText string

// ProtoTmpl returns a sample of a proto file
func ProtoTmpl(out string) error {
	protoFilename := filepath.Base(out)
	serviceName := stringx.From(strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename)))
	text, err := pathx.LoadTemplate(category, rpcTemplateFile, rpcTemplateText)
	if err != nil {
		return err
	}

	dir := filepath.Dir(out)
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

    now := time.Now().Format("2006-01-02 15:04:05")
    winPath := strings.ReplaceAll(filepath.Clean(out), "/", "\\")
    err = util.With("t").Parse(text).SaveTo(map[string]string{
        "package":      serviceName.Untitle(),
        "serviceName":  serviceName.Title(),
        "Date":         now,
        "LastEditTime": now,
        "FilePath":     winPath,
    }, out, false)
    return err
}
