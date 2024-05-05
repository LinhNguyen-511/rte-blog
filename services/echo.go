package services

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ExtractIdFromContext(context echo.Context) (int, error) {
	idParam := context.Param("id")
	return strconv.Atoi(idParam)
}
