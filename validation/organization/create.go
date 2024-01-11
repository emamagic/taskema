package orgvalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)


func (v Validation) CreateOrganization(req param.UserOrganizationCreateRequest) error {
	op := "orgvalidation.CreateOrganization"

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


	return nil
}