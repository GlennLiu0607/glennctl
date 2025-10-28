package generate

import (
    "fmt"

    "github.com/glenn/glennctl/model/tdengine/template"
    "github.com/glenn/glennctl/util/pathx"
)

const (
    category                = "tdengine"
    modelTemplateFile       = "model.tpl"
    modelCustomTemplateFile = "model_custom.tpl"
    modelTypesTemplateFile  = "model_types.tpl"
    errTemplateFile         = "err.tpl"
)

var templates = map[string]string{
    modelTemplateFile:       template.ModelText,
    modelCustomTemplateFile: template.ModelCustomText,
    modelTypesTemplateFile:  template.ModelTypesText,
    errTemplateFile:         template.Error,
}

// Category returns the tdengine category.
func Category() string {
    return category
}

// Clean cleans the tdengine templates.
func Clean() error {
    return pathx.Clean(category)
}

// Templates initializes the tdengine templates.
func Templates() error {
    return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the given template.
func RevertTemplate(name string) error {
    content, ok := templates[name]
    if !ok {
        return fmt.Errorf("%s: no such file name", name)
    }

    return pathx.CreateTemplate(category, name, content)
}

// Update cleans and updates the templates.
func Update() error {
    if err := Clean(); err != nil {
        return err
    }
    return pathx.InitTemplates(category, templates)
}