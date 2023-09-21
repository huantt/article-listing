package forem

import (
	"context"
	"fmt"
	"github.com/huantt/article-listing/model"
	"github.com/huantt/article-listing/pkg/forem"
)

type Service struct {
	foremService *forem.Service
}

func NewService(foremService *forem.Service) *Service {
	return &Service{foremService: foremService}
}

func (s *Service) GetArticles(ctx context.Context, username string, page, perPage int) ([]model.Article, error) {
	articles, err := s.foremService.GetArticles(ctx, forem.GetArticlesPrams{
		UserName: username,
		Page:     page,
		PerPage:  perPage,
	})
	if err != nil {
		return nil, err
	}

	return toModels(articles), nil
}
func toModels(articles []forem.Article) []model.Article {
	var result []model.Article
	for _, article := range articles {
		result = append(result, toModel(article))
	}
	return result
}

func toModel(article forem.Article) model.Article {
	return model.Article{
		Id:          fmt.Sprintf("%d", article.Id),
		Url:         article.Url,
		Title:       article.Title,
		Description: article.Description,
		Thumbnail:   article.CoverImage,
		Author:      article.User.Name,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.EditedAt,
		Tags:        article.TagList,
	}
}
