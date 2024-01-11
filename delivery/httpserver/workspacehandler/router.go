package workspacehandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetupWorkspaceRoutes(echo *echo.Echo) {
	group := echo.Group("/api/v1/users")

	group.POST("/workspace", h.createWorkspace, middleware.Auth(h.authSvc, h.authCfg))
	group.GET("/workspace", h.getAllWorkspaceByOrganizationID, middleware.Auth(h.authSvc, h.authCfg))
	group.DELETE("/workspace/:workspace_id", h.deleteWorkspaceByID, middleware.Auth(h.authSvc, h.authCfg))
}
