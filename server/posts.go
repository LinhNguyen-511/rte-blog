package server

import (
	"log"
	"net/http"
	"rte-blog/services"
	"rte-blog/templates"
	"rte-blog/types"
	"time"

	"github.com/labstack/echo/v4"
)

func (server *server) handleCreatePost(context echo.Context) error {
	title := context.FormValue("title")
	_, err := server.postModel.Create(title)

	return err
}

func (server *server) handleGetPost(context echo.Context) error {
	id, err := services.ExtractIdFromContext(context)
	if err != nil {
		return err
	}

	title, err := server.postModel.GetById(id)
	post := types.Post{
		Id:          id,
		Title:       title,
		AuthorName:  "",
		PublishedAt: time.Now(),
		Content:     []string{title},
	}

	if err != nil {
		log.Fatal(err)
	}

	return templates.Render(context, http.StatusOK, templates.PostLayout(post))
}

func (server *server) handlePutPostTitle(context echo.Context) error {
	id, err := services.ExtractIdFromContext(context)
	if err != nil {
		return err
	}

	title := context.FormValue("title")
	post := types.Post{
		Title: title,
		Id:    id,
	}

	_, err = server.postModel.PutTitle(post)

	return err
}
