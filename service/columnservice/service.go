package columnservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

type Repository interface {
	CreateColumn(column entity.Column) (uint, error)
	GetAllColumnByWorkspaceID(workspaceID uint) ([]entity.Column, error)
	DeleteColumnByID(columnID uint) error
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

func (s Service) CreateColumn(req param.UserColumnCreateRequest) (uint, error) {
	op := "columnservice.CreateColumn"

	column := entity.Column{
		Title:         req.Title,
		CreatorUserID: req.CreatorUserID,
		WorkspaceID:   req.WorkspaceID,
	}

	id, err := s.repo.CreateColumn(column)

	if err != nil {

		return 0, richerror.New(op).WithError(err)
	}

	return id, nil
}

func (s Service) GetAllColumnByWorkspaceID(req param.UserColumnGetAllRequest) ([]param.UserColumnGetAllResponse, error) {
	op := "columnservice.GetAllColumnByWorkspaceID"

	columns, err := s.repo.GetAllColumnByWorkspaceID(req.WorkspaceID)
	if err != nil {

		return nil,
			richerror.New(op).WithError(err)
	}
	return param.ColumnFromEntities(columns), nil
}

func (s Service) DeleteColumnByID(req param.UserColumnDeleteRequest) error {
	op := "columnservice.DeleteColumnByID"

	if err := s.repo.DeleteColumnByID(req.ColumnID); err != nil {

		return richerror.New(op).WithError(err)
	}

	return nil
}
