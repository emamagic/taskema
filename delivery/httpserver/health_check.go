package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})
}
