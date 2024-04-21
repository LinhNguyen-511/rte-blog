package server

import (
	"net/http"
	"rte-blog/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

func home(context echo.Context) error {
	return templates.Render(context, http.StatusOK, templates.CreatePostButton())
}

func (server *server) handleCreatePost(context echo.Context) error {
	title := context.FormValue("title")
	_, err := server.postModel.Create(title)

	return err
}

func (server *server) getPost(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	title, err := server.postModel.GetById(id)

	if err != nil {
		return err
	}

	return templates.Render(context, http.StatusOK, templates.Post(title))
}
