package filehandler

import (
	"net/http"
	"taskema/pkg/richerror"

	"github.com/labstack/echo/v4"
)

func (h Handler) downloadFile(c echo.Context) error {	

	hash := c.QueryParam("hash")
	if hash == "" {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "no hash passed",
		})
	}

	path, err := h.fileSvc.GetFile(hash)
	if err != nil {
		cErr := err.(richerror.RichError)
		if cErr.Code() == richerror.CodeNotFound {
			return echo.NewHTTPError(http.StatusNotFound, echo.Map {
				"message": err.Error(),
			})
		}
	}
	return c.File(path)
}
