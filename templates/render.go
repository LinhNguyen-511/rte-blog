package templates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(context echo.Context, statusCode int, component templ.Component) error {
	context.Response().Writer.WriteHeader(statusCode)
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return component.Render(context.Request().Context(), context.Response().Writer)
}
