package template

import _ "embed"

//go:embed model.tpl
var ModelText string

//go:embed model_custom.tpl
var ModelCustomText string

//go:embed types.tpl
var ModelTypesText string

//go:embed error.tpl
var Error string
