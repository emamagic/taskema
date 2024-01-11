package orgrepo

import "taskema/entity"

type Organization interface {
	CreateOrganization(org entity.Organization) (uint, error)
	GetAllOrganizationByUserID(userID uint) ([]entity.Organization, error)
	DeleteOrganizationByID(orgID uint) error
	GetOrganizationByID(orgID uint) (entity.Organization, error)
}
