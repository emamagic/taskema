package workspacehandler

import (
	"taskema/service/authservice"
	"taskema/service/workspaceservice"
	"taskema/validation/workspacevalidation"
)

type Handler struct {
	authSvc             authservice.Service
	authCfg             authservice.Config
	workspaceservice    workspaceservice.Service
	workspaceValidation workspacevalidation.Validation
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	workspaceservice workspaceservice.Service,
	workspaceValidation workspacevalidation.Validation,
) Handler {
	return Handler{
		authSvc:             authSvc,
		authCfg:             authCfg,
		workspaceservice:    workspaceservice,
		workspaceValidation: workspaceValidation,
	}
}
