package server

import (
	"net/http"
	"rte-blog/data"
	"rte-blog/templates"

	"github.com/labstack/echo/v4"
)

func (model *model) GetIndex(context echo.Context) error {
	return templates.Render(context, http.StatusOK, templates.CreatePostButton())
}

func (model *model) CreatePost(context echo.Context) error {
	_, err := data.CreatePost(model.db, "")

	return err
}
