package forem

import (
	_ "embed"
	"encoding/json"
	"github.com/huantt/article-listing/pkg/forem"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testdata/articles.json
var articleData []byte

func TestToModels(t *testing.T) {
	var articles []forem.Article
	err := json.Unmarshal(articleData, &articles)
	if err != nil {
		panic(err)
	}

	models := toModels(articles)
	assert.Equal(t, len(articles), len(models))
}
