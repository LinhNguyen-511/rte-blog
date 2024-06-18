package data

import (
	"database/sql"
	"rte-blog/types"
)

type PostStore interface {
	Create(title string) (int, error)
	GetById(id int) (post *types.Post, err error)
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

func (model *PostModel) GetById(postId int) (post *types.Post, err error) {
	post = &types.Post{}

	rows, err := model.Store.Query(`
	SELECT  p.title, 
			a.name as author_name, 
			pc.content_id, pc.content_type, pc.order_in_post, 
  		CASE pc.content_type
    		WHEN 'paragraphs' THEN pa.value
    	ELSE NULL
  		END AS value
	FROM posts p
	LEFT JOIN authors a ON a.id = p.author_id
	LEFT JOIN posts_contents pc ON pc.post_id = p.id
	LEFT JOIN paragraphs pa ON pc.content_id = pa.content_id
	WHERE p.id = $1;
	`, postId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// iterate the rows and populate the post's contents
	for rows.Next() {
		var authorName sql.NullString
		var title, contentType, contentValue string
		var contentId, orderInPost int

		if err := rows.Scan(&title, &authorName, &contentId, &contentType, &orderInPost, &contentValue); err != nil {
			return nil, err
		}

		if post.Title == "" || post.Id == 0 {
			post.Title = title
			post.Id = postId
		}

		if post.AuthorName == "" {
			if authorName.Valid {
				post.AuthorName = authorName.String
			} else {
				post.AuthorName = ""
			}
		}

		content := types.Content{
			ContentId: contentId,
			Type:      contentType,
			Value:     contentValue,
		}

		post.Contents = append(post.Contents, content)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return post, nil
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

	if err = transaction.Commit(); err != nil {
		return nil, err
	}

	return paragraph, nil
}

// func (model *PostModel) GetAllContents() ([]types.Content, error) {
// 	// transaction to return all contents in order

// 	return []types.Content{}, nil
// }
