package orghandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetupOrgRoutes(echo *echo.Echo) {
	group := echo.Group("/api/v1/users")

	group.POST("/organization", h.createOrganization, middleware.Auth(h.authSvc, h.authCfg))
	group.GET("/organization", h.getAllOrganization, middleware.Auth(h.authSvc, h.authCfg))
	group.DELETE("/organization/:organization_id", h.deleteOrganizationByID, middleware.Auth(h.authSvc, h.authCfg))
}
