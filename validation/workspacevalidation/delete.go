package workspacevalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"
)

// TODO - avatar existence use for every api, we shoude write it once
func (v Validation) DeleteWorkspace(req param.UserWorkspaceDeleteRequest) error {
	op := "orgvalidation.DeleteWorkspace"

	isExist, oErr := v.workspaceRepo.DoesWorkspaceExist(req.WorkspaceID)
	if oErr != nil {
		return richerror.New(op).WithError(oErr)
	}

	if !isExist {

		return richerror.New(op).WithMessage(richerror.MsgErrorNoWorkspace)
	}

	return nil
}
