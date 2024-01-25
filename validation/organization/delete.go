package orgvalidation

import (
	"errors"
	"taskema/param"
	"taskema/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validation) DeleteOrganization(req param.UserOrganizationDeleteRequest) error {
	op := "orgvalidation.DeleteOrganization"

	vErr := validation.ValidateStruct(&req,
		validation.Field(&req.OrganizationID, validation.Required),
	)

	if vErr != nil {

		return richerror.New(op).WithError(vErr)
	}

	orgExist, err := v.orgRepo.DoesOrganizationExist(req.OrganizationID)
	if err != nil {

		return richerror.New(op).WithError(err)
	}

	if !orgExist {

		return richerror.New(op).WithError(errors.New(richerror.MsgErrorNoOrganization))
	}


	return nil
}