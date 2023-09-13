package collector

import _ "embed"

//go:embed templates/list.tpl
var articleListTemplate string

var templates = []string{
	articleListTemplate,
}
