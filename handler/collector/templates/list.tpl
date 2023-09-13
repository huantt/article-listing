{{- define "article-list"}}
    {{- range $i, $article := .}}
- [{{ truncateByWords $article.Title 10 }}]({{ $article.Url }}) - {{ formatDate $article.CreatedAt }}
    {{- end}}
{{- end}}