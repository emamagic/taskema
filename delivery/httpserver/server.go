package httpserver

import (
	"fmt"
	"taskema/delivery/httpserver/authhandler"
	"taskema/delivery/httpserver/boardhandler"
	"taskema/delivery/httpserver/filehandler"
	"taskema/delivery/httpserver/orghandler"
	"taskema/delivery/httpserver/taskhandler"
	"taskema/delivery/httpserver/userhandler"
	"taskema/delivery/httpserver/workspacehandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Host          string `koanf:"host"`
	Port          uint   `koanf:"port"`
	MaxUploadSize string `koanf:"max_upload_size"`
}

type Server struct {
	cfg              Config
	router           *echo.Echo
	userHandler      userhandler.Handler
	fileHandler      filehandler.Handler
	authHandler      authhandler.Handler
	orgHandler       orghandler.Handler
	workspaceHandler workspacehandler.Handler
	boardHandler     boardhandler.Handler
	taskhandler      taskhandler.Handler
}

func New(
	cfg Config,
	userHandler userhandler.Handler,
	fileHandler filehandler.Handler,
	authHandler authhandler.Handler,
	orgHandler orghandler.Handler,
	workspaceHandler workspacehandler.Handler,
	boardHandler boardhandler.Handler,
	taskhandler taskhandler.Handler,
) Server {
	return Server{
		cfg:              cfg,
		router:           echo.New(),
		userHandler:      userHandler,
		fileHandler:      fileHandler,
		authHandler:      authHandler,
		orgHandler:       orgHandler,
		workspaceHandler: workspaceHandler,
		boardHandler:     boardHandler,
		taskhandler:      taskhandler,
	}
}

func (s Server) Serve() error {

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.BodyLimit(s.cfg.MaxUploadSize))
	s.router.Use(middleware.CORS())

	// TODO - Embeding static file to go executable (goembed or statik)
	s.router.Static("/api-docs", "./delivery/httpserver/apidocs")

	s.router.GET("/health-check", healthCheck)
	s.router.GET("/current-time", currentUnixTime)

	s.userHandler.SetUserRoutes(s.router)
	s.fileHandler.SetFileRouters(s.router)
	s.authHandler.SetAuthRouters(s.router)
	s.orgHandler.SetupOrgRoutes(s.router)
	s.workspaceHandler.SetupWorkspaceRoutes(s.router)
	s.boardHandler.SetupBoardRoutes(s.router)
	s.taskhandler.SetupTaskRoutes(s.router)

	return s.router.StartTLS(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port), "/etc/taskema.ir/ssl/fullchain.pem", "/etc/taskema.ir/ssl/privkey.pem")
}

func (s Server) Router() *echo.Echo {
	return s.router
}
