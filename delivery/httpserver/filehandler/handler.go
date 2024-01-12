package filehandler

import (
	"taskema/service/authservice"
	"taskema/service/fileservice"
)

type Handler struct {
	authSvc authservice.Service
	authCfg authservice.Config
	fileSvc fileservice.Service
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	fileSvc fileservice.Service,
) Handler {
	return Handler{
		authSvc: authSvc,
		authCfg: authCfg,
		fileSvc: fileSvc,
	}
}
