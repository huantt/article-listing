package collector

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/huantt/article-listing/model"
	"html/template"
	"log/slog"
	"os"
	"time"
)

type Handler struct {
	articleService ArticleService
}

func NewHandler(articleService ArticleService) *Handler {
	return &Handler{articleService}
}

func (h *Handler) Collect(ctx context.Context, username string, limit int, templateFilePath, outputFilePath string) error {
	articles, err := h.articleService.GetArticles(ctx, username, 1, limit)
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("Got %d article from %s", len(articles), username))

	templateStr, err := os.ReadFile(templateFilePath)
	if err != nil {
		return err
	}
	output, err := generateOutput(articles, string(templateStr), templates...)
	if err != nil {
		return err
	}
	return os.WriteFile(outputFilePath, []byte(*output), 0644)
}

func generateOutput(articles []model.Article, readmeTemplate string, templates ...string) (*string, error) {
	if len(articles) == 0 {
		return nil, errors.New("articles must be not empty")
	}
	tmpl, err := template.
		New("article").
		Funcs(template.FuncMap{
			"formatDate":      formatDate,
			"formatHour":      formatHour,
			"formatTime":      formatTime,
			"truncateByWords": truncateByWords,
		}).
		Parse(readmeTemplate)
	if err != nil {
		return nil, err
	}

	for _, t := range templates {
		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return nil, err
		}
	}

	var result bytes.Buffer
	err = tmpl.ExecuteTemplate(&result, "article", map[string]any{
		"Articles": articles,
		"Author":   articles[0].Author,
		"Time":     time.Now(),
	})
	if err != nil {
		return nil, err
	}
	stringResult := result.String()
	return &stringResult, nil
}
