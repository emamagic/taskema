package columnhandler

import (
	"taskema/service/authservice"
	"taskema/service/columnservice"
)

type Handler struct {
	authSvc   authservice.Service
	authCfg   authservice.Config
	columnSvc columnservice.Service
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	columnSvc columnservice.Service,
) Handler {
	return Handler{
		authSvc:   authSvc,
		authCfg:   authCfg,
		columnSvc: columnSvc,
	}
}
