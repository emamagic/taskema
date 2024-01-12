package userhandler

import (
	"net/http"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) getUserProfile(c echo.Context) error {

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)

	req := param.UserProfileRequest{UserID: claims.UserID}

	resp, err := h.userSvc.GetUserProfile(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
