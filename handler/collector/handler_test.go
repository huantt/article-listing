package collector

import (
	_ "embed"
	"encoding/json"
	"github.com/huantt/article-listing/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//go:embed testdata/articles.json
var articleData []byte

//go:embed testdata/template.md.tpl
var templateData string

func TestGenerateOutput(t *testing.T) {
	var articles []model.Article
	err := json.Unmarshal(articleData, &articles)
	if err != nil {
		panic(err)
	}

	output, err := generateOutput(articles, templateData, templates...)
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, *output)
	_ = os.WriteFile("../../.test/README.md", []byte(*output), 0644)
}
