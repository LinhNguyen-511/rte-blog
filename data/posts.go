package data

import (
	"database/sql"
	"rte-blog/types"
)

type PostStore interface {
	Create(title string) (int, error)
	GetById(id int) (title string, err error)
	PutTitle(post types.Post) (types.Post, error)
	PostContent(id int, orderInPost int) (types.Content, error)
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

func (model *PostModel) PutTitle(post types.Post) (types.Post, error) {
	_, err := model.Store.Exec("UPDATE posts SET title = $1 WHERE id = $2", post.Title, post.Id)
	return post, err
}

func (model *PostModel) PostContent(id int, orderInPost int) (types.Content, error) {
	// insert into posts_contents -> content_id

	// insert into paragraphs

	return types.Content{}, nil
}

// func (model *PostModel) GetAllContents() ([]types.Content, error) {
// 	// transaction to return all contents in order

// 	return []types.Content{}, nil
// }
