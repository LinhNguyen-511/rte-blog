package data

import (
	"database/sql"
	"fmt"
	"rte-blog/types"
)

type PostStore interface {
	Create(title string) (int, error)
	GetById(id int) (title string, err error)
	PutTitle(post types.Post) (types.Post, error)
	PostContent(id int) (types.Content, error)
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

func (model *PostModel) PostContent(id int) (types.Content, error) {
	var contentId int
	err := model.Store.QueryRow("INSERT INTO paragraphs(value) VALUES($1) RETURNING id", "").Scan(&contentId)
	model.Store.QueryRow("UPDATE posts SET contents = array_append(contents, $1) WHERE id = $2", fmt.Sprintf("paragraphs:%d", contentId), id)

	return types.Content{Id: contentId, Value: "", Type: types.ContentParagraph}, err
}
