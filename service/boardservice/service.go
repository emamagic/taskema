package boardservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

type Repository interface {
	CreateBoard(board entity.Board) (uint, error)
	GetAllBoardByWorkspaceID(workspaceID uint) ([]entity.Board, error)
	DeleteBoardByID(boardID uint) error

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

func (s Service) CreateBoard(req param.UserBoardCreateRequest) (uint, error) {
	op := "boardservice.CreateBoard"

	board := entity.Board{
		Title:         req.Title,
		Avatar:        req.Avatar,
		CreatorUserID: req.CreatorUserID,
		WorkspaceID:   req.WorkspaceID,
	}

	id, err := s.repo.CreateBoard(board)

	if err != nil {

		return 0, richerror.New(op).WithError(err)
	}

	return id, nil
}

func (s Service) GetAllBoardByWorkspaceID(req param.UserBoardGetAllRequest) ([]param.UserBoardGetAllResponse, error) {
	op := "boardservice.GetAllBoardByWorkspaceID"

	boards, err := s.repo.GetAllBoardByWorkspaceID(req.WorkspaceID)
	if err != nil {

		return nil,
			richerror.New(op).WithError(err)
	}
	return param.BoardFromEntities(boards), nil
}

func (s Service) DeleteBoardByID(req param.UserBoardDeleteRequest) error {
	op := "boardservice.DeleteBoardByID"

	if err := s.repo.DeleteBoardByID(req.BoardID); err != nil {

		return richerror.New(op).WithError(err)
	}

	return nil
}
