package server

import (
	"net/http"
	"rte-blog/templates"

	"github.com/labstack/echo/v4"
)

func (model *model) GetPost(context echo.Context) error {
	return templates.Render(context, http.StatusOK, templates.CreatePostButton())
}
