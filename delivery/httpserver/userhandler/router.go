package userhandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	group := e.Group("/api/v1/users")

	group.POST("/register", h.register)
	group.POST("/login", h.login)
	group.GET("/profile", h.getUserProfile, middleware.Auth(h.authSvc, h.authCfg))

}
