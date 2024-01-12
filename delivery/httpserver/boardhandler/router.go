package boardhandler

import (
	"taskema/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetupBoardRoutes(echo *echo.Echo) {
	group := echo.Group("/api/v1/users")

	group.POST("/board", h.createBoard, middleware.Auth(h.authSvc, h.authCfg))
	group.GET("/board", h.getAllBoardByWorkspaceID, middleware.Auth(h.authSvc, h.authCfg))
	group.DELETE("/board/:board_id", h.deleteBoardByID, middleware.Auth(h.authSvc, h.authCfg))
}
