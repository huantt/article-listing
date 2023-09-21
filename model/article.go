package model

import "time"

type Article struct {
	Id          string     `json:"id"`
	Url         string     `json:"url"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Thumbnail   string     `json:"thumbnail"`
	Author      string     `json:"author"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Tags        []string   `json:"tags"`
}
