package param

import (
	"taskema/entity"
)

type UserTaskCreateRequest struct {
	CreatorUserID  uint    `json:"creator_user_id"`
	Title          string  `json:"title"`
	Avatar         *string `json:"avatar"`
	Description    string  `json:"description"`
	ColumnID       uint    `json:"column_id"`
	AssignedUserID uint    `json:"assigned_user_id"`
	DueDate        int64   `json:"due_date"`
}

type UserTaskGetAllRequest struct {
	UserID   uint
	ColumnID uint `json:"column_id"`
}

type UserTaskDeleteRequest struct {
	TaskID uint `json:"task_id"`
}

type UserTaskResponse struct {
	ID             uint    `json:"id"`
	Title          string  `json:"title"`
	Avatar         *string `json:"avatar"`
	CreatorUserID  uint    `json:"creator_user_id"`
	Description    string  `json:"description"`
	DueDate        *int64  `json:"due_date"`
	ColumnID       uint    `json:"column_id"`
	AssignedUserID uint    `json:"assigned_user_id"`
	Priority       uint    `json:"priority"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
}

func TaskFromEntities(entities []entity.Task) []UserTaskResponse {
	list := make([]UserTaskResponse, 0)

	for _, task := range entities {
		resp := UserTaskResponse{
			ID:             task.ID,
			Title:          task.Title,
			Avatar:         task.Avatar,
			CreatorUserID:  task.CreatorUserID,
			Description:    task.Description,
			DueDate:        task.DueDate,
			ColumnID:       task.ColumnID,
			AssignedUserID: task.AssignedUserID,
			Priority:       task.Priority,
			CreateAt:       task.CreateAt.Unix(),
			UpdateAt:       task.UpdateAt.Unix(),
		}
		list = append(list, resp)
	}

	return list
}
