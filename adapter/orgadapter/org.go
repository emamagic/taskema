package orgadapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/orgrepo"
)

type Adapter struct {
	orgRepo orgrepo.Organization
}

func New(
	orgRepo orgrepo.Organization,
) Adapter {
	return Adapter{
		orgRepo: orgRepo,
	}
}

func (a Adapter) DoesOrganizationExist(organizationID uint) (bool, error) {

	_, err := a.orgRepo.GetOrganizationByID(organizationID)
	if err != nil {
		cErr := err.(richerror.RichError)
		if cErr.Code() == richerror.CodeNotFound {

			return false, nil
		} else {

			return false, err
		}
	}

	return true, nil
}
