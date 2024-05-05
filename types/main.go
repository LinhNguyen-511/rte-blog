package types

import "time"

type Post struct {
	Id          int
	Title       string
	AuthorName  string
	PublishedAt time.Time
	Content     []string
}
