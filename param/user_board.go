package param

import (
	"taskema/entity"
	"time"
)

type UserBoardCreateRequest struct {
	CreatorUserID uint    `json:"creator_user_id"`
	Title         string  `json:"title"`
	Avatar        *string `json:"avatar"`
	WorkspaceID   uint    `json:"workspace_id"`
}

type UserBoardGetAllRequest struct {
	UserID      uint `json:"user_id"`
	WorkspaceID uint `json:"workspace_id"`
}

type UserBoardDeleteRequest struct {
	BoardID uint `json:"board_id"`
}

type UserBoardGetAllResponse struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Avatar        *string   `json:"avatar"`
	WorkspaceID   uint      `json:"workspace_id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
}

func BoardFromEntities(entities []entity.Board) []UserBoardGetAllResponse {
	list := make([]UserBoardGetAllResponse, 0)

	for _, board := range entities {
		resp := UserBoardGetAllResponse{
			ID:            board.ID,
			Title:         board.Title,
			Avatar:        board.Avatar,
			WorkspaceID:   board.WorkspaceID,
			CreateAt:      board.CreateAt,
			UpdateAt:      board.UpdateAt,
		}
		list = append(list, resp)
	}

	return list
}
