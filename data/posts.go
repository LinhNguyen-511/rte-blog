package data

import (
	"database/sql"
	"rte-blog/types"
)

type PostStore interface {
	Create(title string) (int, error)
	GetById(id int) (title string, err error)
	PutTitle(post types.Post) (types.Post, error)
	CreatePostContent(id int, orderInPost int) (*types.Content, error)
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

func (model *PostModel) CreatePostContent(postId int, orderInPost int) (*types.Content, error) {
	paragraph := &types.Content{
		Type:  types.ContentParagraphs,
		Value: "",
	}

	// transaction
	transaction, err := model.Store.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback()
	// insert into posts_contents -> content_id
	var contentId int
	var paragraphId int

	err = transaction.QueryRow(`INSERT INTO posts_contents(content_type, post_id, order_in_post) VALUES ($1, $2, $3) RETURNING content_id`, types.ContentParagraphs, postId, orderInPost).Scan(&contentId)
	if err != nil {
		return nil, err
	}
	// insert into paragraphs
	err = transaction.QueryRow(`INSERT INTO paragraphs(content_id, value) VALUES ($1, $2) RETURNING id`, contentId, "").Scan(&paragraphId)
	if err != nil {
		return nil, err
	}

	paragraph.ContentId = contentId
	paragraph.Id = paragraphId

	if err = transaction.Commit(); err != nil {
		return nil, err
	}

	return paragraph, nil
}

// func (model *PostModel) GetAllContents() ([]types.Content, error) {
// 	// transaction to return all contents in order

// 	return []types.Content{}, nil
// }
