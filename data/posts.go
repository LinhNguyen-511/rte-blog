package data

import "database/sql"

func CreatePost(db *sql.DB, title string) (int, error) {
	var postId int
	err := db.QueryRow("INSERT INTO posts(title) VALUES($1) RETURNING id", title).Scan(&postId)

	return postId, err
}

func GetPost(db *sql.DB, id int) (title string) {
	db.QueryRow("SELECT title FROM posts WHERE id = $1", id).Scan(&title)
	return
}
