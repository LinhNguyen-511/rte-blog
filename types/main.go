package types

import "time"

type DbConfig struct {
	DbName   string
	User     string
	Password string
}

type Post struct {
	Id          int
	Title       string
	AuthorName  string
	PublishedAt time.Time
	Contents    []Content
}

type Content struct {
	ContentId   int
	Value       string
	Type        string
	OrderInPost int
}

const (
	ContentParagraphs = "paragraphs"
)
