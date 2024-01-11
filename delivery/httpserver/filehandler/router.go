package filehandler

import (
	"taskema/delivery/httpserver/middleware"
	
	"github.com/labstack/echo/v4"
)

func (h Handler) SetFileRouters(e *echo.Echo) {
	group := e.Group("/api/v1/files")

	group.GET("/download", h.downloadFile, middleware.Auth(h.authSvc, h.authCfg))
	group.POST("/upload", h.uploadFile, middleware.Auth(h.authSvc, h.authCfg))
}
