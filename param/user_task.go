package param

import (
	"taskema/entity"
	"time"
)

type UserTaskCreateRequest struct {
	CreatorUserID uint    `json:"creator_user_id"`
	Title         string  `json:"title"`
	Avatar        *string `json:"avatar"`
	Description   string  `json:"description"`
	BoardID       uint    `json:"board_id"`
}

type UserTaskGetAllRequest struct {
	UserID  uint `json:"user_id"`
	BoardID uint  `json:"board_id"`
}

type UserTaskDeleteRequest struct {
	TaskID uint `json:"task_id"`
}

type UserTaskResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Avatar         *string   `json:"avatar"`
	CreatorUserID  uint      `json:"creator_user_id"`
	Description    string    `json:"description"`
	DueDate        time.Time `json:"due_date"`
	BoardID        uint      `json:"board_id"`
	AssignedUserID uint      `json:"assigned_user_id"`
	Priority       uint      `json:"priority"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
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
			BoardID:        task.BoardID,
			AssignedUserID: task.AssignedUserID,
			Priority:       task.Priority,
			CreateAt:       task.CreateAt,
			UpdateAt:       task.UpdateAt,
		}
		list = append(list, resp)
	}

	return list
}
