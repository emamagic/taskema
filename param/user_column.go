package param

import (
	"taskema/entity"
)

type UserColumnCreateRequest struct {
	Title         string `json:"title"`
	WorkspaceID   uint   `json:"workspace_id"`
	CreatorUserID uint
}

type UserColumnGetAllRequest struct {
	WorkspaceID uint `json:"workspace_id"`
	UserID      uint
}

type UserColumnDeleteRequest struct {
	UserID   uint
	ColumnID uint `json:"column_id"`
}

type UserColumnGetAllResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	WorkspaceID uint   `json:"workspace_id"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func ColumnFromEntities(entities []entity.Column) []UserColumnGetAllResponse {
	list := make([]UserColumnGetAllResponse, 0)

	for _, column := range entities {
		resp := UserColumnGetAllResponse{
			ID:          column.ID,
			Title:       column.Title,
			WorkspaceID: column.WorkspaceID,
			CreateAt:    column.CreateAt.Unix(),
			UpdateAt:    column.UpdateAt.Unix(),
		}
		list = append(list, resp)
	}

	return list
}
