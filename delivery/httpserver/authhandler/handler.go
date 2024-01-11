package authhandler

import "taskema/service/authservice"

type Handler struct {
	authSvc authservice.Service
	authCfg authservice.Config
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
) Handler {
	return Handler{
		authSvc: authSvc,
		authCfg: authCfg,
	}
}
