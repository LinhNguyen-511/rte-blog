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
	Content     []string
}
