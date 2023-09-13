package forem

import (
	"time"
)

type Article struct {
	TypeOf                 string     `json:"type_of"`
	Id                     int        `json:"id"`
	Title                  string     `json:"title"`
	Description            string     `json:"description"`
	ReadablePublishDate    string     `json:"readable_publish_date"`
	Slug                   string     `json:"slug"`
	Path                   string     `json:"path"`
	Url                    string     `json:"url"`
	CommentsCount          int        `json:"comments_count"`
	PublicReactionsCount   int        `json:"public_reactions_count"`
	PublishedTimestamp     *time.Time `json:"published_timestamp"`
	PositiveReactionsCount int        `json:"positive_reactions_count"`
	CoverImage             string     `json:"cover_image"`
	SocialImage            string     `json:"social_image"`
	CanonicalUrl           string     `json:"canonical_url"`
	CreatedAt              *time.Time `json:"created_at"`
	EditedAt               *time.Time `json:"*"`
	PublishedAt            *time.Time `json:"published_at"`
	LastCommentAt          *time.Time `json:"last_comment_at"`
	ReadingTimeMinutes     int        `json:"reading_time_minutes"`
	TagList                []string   `json:"tag_list"`
	Tags                   string     `json:"tags"`
	User                   User       `json:"user"`
}

type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	UserId          int    `json:"user_id"`
	WebsiteUrl      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`
	ProfileImage90  string `json:"profile_image_90"`
}

type GetArticlesPrams struct {
	MostRecent   bool
	Page         int
	PerPage      int
	Tag          string
	Tags         []string
	TagsExclude  []string
	UserName     string
	State        string
	Top          int
	CollectionID int
}
