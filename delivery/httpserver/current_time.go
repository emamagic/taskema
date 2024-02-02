package httpserver

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func currentUnixTime(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"current-time": time.Now().Unix(),
	})
}
