package boardhandler

import (
	"taskema/service/authservice"
	"taskema/service/boardservice"
)

type Handler struct {
	authSvc  authservice.Service
	authCfg  authservice.Config
	boardSvc boardservice.Service
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	boardSvc boardservice.Service,
) Handler {
	return Handler{
		authSvc: authSvc,
		authCfg: authCfg,
		boardSvc: boardSvc,
	}
}
