package data

import "database/sql"

type PostStore interface {
	Create(title string) (int, error)
	GetById(id int) (title string, err error)
	PutTitle(title string, id int) error
}

type PostModel struct {
	Store *sql.DB
}

func NewPostModel(store *sql.DB) *PostModel {
	return &PostModel{store}
}

func (model *PostModel) Create(title string) (postId int, err error) {
	err = model.Store.QueryRow("INSERT INTO posts(title) VALUES($1) RETURNING id", title).Scan(&postId)
	return postId, err
}

func (model *PostModel) GetById(id int) (title string, err error) {
	model.Store.QueryRow("SELECT title FROM posts WHERE id = $1", id).Scan(&title)
	return title, nil
}

func (model *PostModel) PutTitle(title string, id int) (err error) {
	_, err = model.Store.Exec("UPDATE posts SET title = $1 WHERE id = $2", title, id)
	return err
}