package workspacevalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// TODO - avatar existence use for every api, we shoude write it once
func (v Validation) CreateWorkspace(req param.UserWorkspaceCreateRequest) error {
	op := "orgvalidation.CreateWorkspace"

	vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required, is.Alpha, validation.Length(3, 10)),
	)

	if vErr != nil {

		return richerror.New(op).WithError(vErr)
	}

	if req.Avatar != nil {

		isAvatarExist, aErr := v.fileRepo.DoesFileExist(*req.Avatar)
		if aErr != nil {
			return richerror.New(op).WithError(aErr)
		}

		if !isAvatarExist {
			return richerror.New(op).
				WithMessage(richerror.MsgErrorAvatarNotValid).
				WithCode(richerror.CodeNotFound)
		}

	}

	isOrganizationExist, oErr := v.orgRepo.DoesOrganizationExist(req.OrganizationID)
	if oErr != nil {
		return richerror.New(op).WithError(oErr)
	}

	if !isOrganizationExist {

		return richerror.New(op).WithMessage(richerror.MsgErrorNoOrganization)
	}

	return nil
}
