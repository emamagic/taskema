package userhandler

import (
	"net/http"
	"taskema/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) register(c echo.Context) error {
	var req param.UserRegisterRequest
	
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	vErr := h.userValidator.Register(req)
	if vErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": vErr.Error(),
		})
	}

	resp, rErr := h.userSvc.Register(req)
	if rErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": rErr.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
