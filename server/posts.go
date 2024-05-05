package server

import (
	"log"
	"net/http"
	"rte-blog/templates"
	"rte-blog/types"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (server *server) handleCreatePost(context echo.Context) error {
	title := context.FormValue("title")
	_, err := server.postModel.Create(title)

	return err
}

func (server *server) handleGetPost(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	title, err := server.postModel.GetById(id)
	post := types.Post{
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
