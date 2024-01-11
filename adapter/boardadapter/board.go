package boardadapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/boardrepo"
)

type Adapter struct {
	boardRepo boardrepo.Board
}

func New(
	boardRepo boardrepo.Board,
) Adapter {
	return Adapter{
		boardRepo: boardRepo,
	}
}

func (a Adapter) DoesBoardExist(boardID uint) (bool, error) {

	_, err := a.boardRepo.GetBoardByID(boardID)
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
