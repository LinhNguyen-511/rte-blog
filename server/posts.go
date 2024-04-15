package server

import (
	"net/http"
	"rte-blog/data"
	"rte-blog/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

type postModel struct {
	model
}

func (model *postModel) getIndex(context echo.Context) error {
	return templates.Render(context, http.StatusOK, templates.CreatePostButton())
}

func (model *postModel) create(context echo.Context) error {
	_, err := data.CreatePost(model.db, "")

	return err
}

func (model *postModel) get(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}

	title := data.GetPost(model.db, id)

	return templates.Render(context, http.StatusOK, templates.Post(title))
}
