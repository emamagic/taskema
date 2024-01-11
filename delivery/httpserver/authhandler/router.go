package authhandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetAuthRouters(e *echo.Echo) {
	group := e.Group("/api/v1/auth")

	group.POST("/refresh", h.refreshToken, middleware.Auth(h.authSvc, h.authCfg))
}