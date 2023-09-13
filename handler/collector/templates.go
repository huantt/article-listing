package collector

import _ "embed"

//go:embed templates/list.tpl
var articleListTemplate string

//go:embed templates/table.tpl
var articleTableTemplate string

var templates = []string{
	articleListTemplate, articleTableTemplate,
}
