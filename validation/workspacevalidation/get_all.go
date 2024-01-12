package workspacevalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"
)

func (v Validation) GetAllWorkspace(req param.UserWorkspaceGetAllRequest) error {
	op := "orgvalidation.GetAllWorkspace"

	isOrganizationExist, oErr := v.orgRepo.DoesOrganizationExist(req.OrganizationID)
	if oErr != nil {

		return richerror.New(op).WithError(oErr)
	}

	if !isOrganizationExist {
		return richerror.New(op).WithMessage(richerror.MsgErrorNoOrganization)
	}

	return nil
}