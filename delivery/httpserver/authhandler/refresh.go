package authhandler

import (
	"net/http"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) refreshToken(c echo.Context) error {

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)

	accessToken, aErr := h.authSvc.GenerateAccessToken(claims.UserID, claims.RoleID)
	if aErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": aErr.Error(),
		})
	}

	refreshToken, rErr := h.authSvc.GenerateRefreshToken(claims.UserID, claims.RoleID)
	if rErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": rErr.Error(),
		})
	}

	resp := param.AuthRefreshtokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, resp)
}
