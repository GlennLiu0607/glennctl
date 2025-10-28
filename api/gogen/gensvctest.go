package gogen

import (
	_ "embed"

	"github.com/GlennLiu0607/glennctl/api/spec"
	"github.com/GlennLiu0607/glennctl/config"
	"github.com/GlennLiu0607/glennctl/internal/version"
	"github.com/GlennLiu0607/glennctl/util/format"
)

//go:embed svc_test.tpl
var svcTestTemplate string

func genServiceContextTest(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + "_test.go",
		templateName:    "svcTestTemplate",
		category:        category,
		templateFile:    svcTestTemplateFile,
		builtinTemplate: svcTestTemplate,
		data: map[string]any{
			"projectPkg": projectPkg,
			"version":    version.BuildVersion,
		},
	})
}
