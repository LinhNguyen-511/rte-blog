package types

import "time"

type Post struct {
	Title       string
	AuthorName  string
	PublishedAt time.Time
	Content     []string
}
