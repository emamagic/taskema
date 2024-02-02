package columnhandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetupColumnRoutes(echo *echo.Echo) {
	group := echo.Group("/api/v1/users")

	group.POST("/column", h.createColumn, middleware.Auth(h.authSvc, h.authCfg))
	group.GET("/column", h.getAllColumnByWorkspaceID, middleware.Auth(h.authSvc, h.authCfg))
	group.DELETE("/column/:column_id", h.deleteColumnByID, middleware.Auth(h.authSvc, h.authCfg))
}
