package taskhandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetupTaskRoutes(echo *echo.Echo) {
	group := echo.Group("/api/v1/users")

	group.POST("/task", h.createTask, middleware.Auth(h.authSvc, h.authCfg))
	group.GET("/task", h.getAllTaskByColumnID, middleware.Auth(h.authSvc, h.authCfg))
	group.DELETE("/task/:task_id", h.deleteTaskByID, middleware.Auth(h.authSvc, h.authCfg))
}
