package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Print(c echo.Context) error {
	return c.JSON(http.StatusOK, "printing hello world")
}
