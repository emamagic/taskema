package param

import (
	"taskema/entity"
)

type UserBoardCreateRequest struct {
	CreatorUserID uint
	Title         string  `json:"title"`
	Avatar        *string `json:"avatar"`
	WorkspaceID   uint    `json:"workspace_id"`
}

type UserBoardGetAllRequest struct {
	UserID      uint
	WorkspaceID uint `json:"workspace_id"`
}

type UserBoardDeleteRequest struct {
	UserID  uint
	BoardID uint `json:"board_id"`
}

type UserBoardGetAllResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Avatar      *string `json:"avatar"`
	WorkspaceID uint    `json:"workspace_id"`
	CreateAt    int64   `json:"create_at"`
	UpdateAt    int64   `json:"update_at"`
}

func BoardFromEntities(entities []entity.Board) []UserBoardGetAllResponse {
	list := make([]UserBoardGetAllResponse, 0)

	for _, board := range entities {
		resp := UserBoardGetAllResponse{
			ID:          board.ID,
			Title:       board.Title,
			Avatar:      board.Avatar,
			WorkspaceID: board.WorkspaceID,
			CreateAt:    board.CreateAt.Unix(),
			UpdateAt:    board.UpdateAt.Unix(),
		}
		list = append(list, resp)
	}

	return list
}
