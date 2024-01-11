package userhandler

import (
	"taskema/service/authservice"
	"taskema/service/userservice"
	"taskema/validation/uservalidation"
)

type Handler struct {
	userSvc       userservice.Service
	userValidator uservalidation.Validator
	authSvc       authservice.Service
	authCfg       authservice.Config
}

func New(
	userSvc userservice.Service,
	userValidator uservalidation.Validator,
	authSvc authservice.Service,
	authCfg authservice.Config,
) Handler {
	return Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
		authSvc:       authSvc,
		authCfg:       authCfg,
	}
}
