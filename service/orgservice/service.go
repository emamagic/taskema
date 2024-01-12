package orgservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

type Repository interface {
	CreateOrganization(org entity.Organization) (uint, error)
	GetAllOrganizationByUserID(userID uint) ([]entity.Organization, error)
	DeleteOrganizationByID(orgID uint) error
}

type Service struct {
	repo Repository
}

func New(
	repo Repository,
) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateOrganization(req param.UserOrganizationCreateRequest) (uint, error) {
	op := "orgservice.CreateOrganization"

	org := entity.Organization{
		Title:         req.Title,
		Avatar:        req.Avatar,
		CreatorUserID: req.CreatorUserID,
	}

	id, err := s.repo.CreateOrganization(org)

	if err != nil {

		return 0, richerror.New(op).WithError(err)
	}

	return id, nil
}

func (s Service) GetAllOrganizationByID(req param.UserOrganizationGetAllRequest) ([]param.UserOrganizationResponse, error) {
	op := "orgservice.GetAllOrganizationByID"

	users, err := s.repo.GetAllOrganizationByUserID(req.UserID)
	if err != nil {

		return nil,
			richerror.New(op).WithError(err)
	}

	return param.OrganizationFromEntities(users), nil
}

func (s Service) DeleteOrganizationByID(req param.UserOrganizationDeleteRequest) error {
	op := "orgservice.DeleteOrganizationByID"

	if err := s.repo.DeleteOrganizationByID(req.OrganizationID); err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return nil
}
