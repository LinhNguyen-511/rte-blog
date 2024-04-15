package data

import "database/sql"

func CreatePost(db *sql.DB, title string) (int, error) {
	var postId int
	err := db.QueryRow("INSERT INTO posts(title) VALUES($1) RETURNING id", title).Scan(&postId)

	return postId, err
}
