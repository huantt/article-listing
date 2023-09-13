package forem

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetArticles(t *testing.T) {
	userName := "jacktt"
	perPage := 10
	service := NewService(DevToEndpoint)
	articles, err := service.GetArticles(context.Background(), GetArticlesPrams{
		UserName: userName,
		PerPage:  perPage,
	})
	assert.NoError(t, err)
	assert.Equal(t, 10, len(articles))
	for _, article := range articles {
		assert.Equal(t, userName, article.User.Username)
	}
}
