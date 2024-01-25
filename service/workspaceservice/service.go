package workspaceservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

type Repository interface {
	CreateWorkspace(workspace entity.Workspace) (uint, error)
	GetAllWorkspaceByOrganizationID(organizationID uint, userID uint) ([]entity.Workspace, error)
	DeleteWorkspaceByID(workspaceID uint) error
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

func (s Service) CreateWorkspace(req param.UserWorkspaceCreateRequest) (uint, error) {
	op := "workspaceservice.CreateWorkspace"

	workspace := entity.Workspace{
		Title:          req.Title,
		Avatar:         req.Avatar,
		CreatorUserID:  req.CreatorUserID,
		OrganizationID: req.OrganizationID,
	}

	id, err := s.repo.CreateWorkspace(workspace)

	if err != nil {

		return 0, richerror.New(op).WithError(err)
	}

	return id, nil
}

func (s Service) GetAllWorkspaceByOrganizationID(req param.UserWorkspaceGetAllRequest) ([]param.UserWorkspaceResponse, error) {
	op := "workspaceservice.GetAllWorkspaceByOrganizationID"

	workspaces, err := s.repo.GetAllWorkspaceByOrganizationID(req.OrganizationID, req.UserID)
	if err != nil {

		return nil,
			richerror.New(op).WithError(err)
	}
	return param.WorkspaceFromEntities(workspaces), nil

}

func (s Service) DeleteWorkspaceByID(req param.UserWorkspaceDeleteRequest) error {
	op := "workspaceservice.DeleteWorkspaceByID"

	if err := s.repo.DeleteWorkspaceByID(req.WorkspaceID); err != nil {

		return richerror.New(op).WithError(err)
	}

	return nil
}
