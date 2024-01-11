package taskhandler

import (
	"taskema/service/authservice"
	"taskema/service/taskservice"
)

type Handler struct {
	authSvc     authservice.Service
	authCfg     authservice.Config
	taskservice taskservice.Service
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	taskservice taskservice.Service,
) Handler {
	return Handler{
		authSvc:     authSvc,
		authCfg:     authCfg,
		taskservice: taskservice,
	}
}
