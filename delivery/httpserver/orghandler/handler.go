package orghandler

import (
	"taskema/service/authservice"
	"taskema/service/orgservice"
	orgvalidation "taskema/validation/organization"
)

type Handler struct {
	authSvc       authservice.Service
	authCfg       authservice.Config
	orgSvc        orgservice.Service
	orgValidation orgvalidation.Validation
}

func New(
	authSvc authservice.Service,
	authCfg authservice.Config,
	orgSvc orgservice.Service,
	orgValidation orgvalidation.Validation,
) Handler {
	return Handler{
		authSvc:       authSvc,
		authCfg:       authCfg,
		orgSvc:        orgSvc,
		orgValidation: orgValidation,
	}
}
