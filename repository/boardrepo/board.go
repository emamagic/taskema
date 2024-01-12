package boardrepo

import "taskema/entity"

type Board interface {
	CreateBoard(board entity.Board) (uint, error)
	GetAllBoardByWorkspaceID(workspaceID uint) ([]entity.Board, error)
	DeleteBoardByID(boardID uint) error
	GetBoardByID(boardID uint) (entity.Board, error)
}
