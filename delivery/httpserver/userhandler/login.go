package userhandler

import (
	"net/http"
	"taskema/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) login(c echo.Context) error {
	var req param.UserLoginRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	vErr := h.userValidator.Login(req)
	if vErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": vErr.Error(),
		})
	}

	resp, lErr := h.userSvc.Login(req)
	if lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": lErr.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
