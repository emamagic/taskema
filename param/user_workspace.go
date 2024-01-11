package param

import (
	"taskema/entity"
	"time"
)

type UserWorkspaceCreateRequest struct {
	CreatorUserID  uint    `json:"creator_user_id"`
	Title          string  `json:"title"`
	Avatar         *string `json:"avatar"`
	OrganizationID uint    `json:"organization_id"`
}

type UserWorkspaceGetAllRequest struct {
	UserID uint `json:"user_id"`
	OrganizationID uint `json:"organization_id"`
}

type UserWorkspaceDeleteRequest struct {
	WorkspaceID uint `json:"workspace_id"`
}

type UserWorkspaceResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Avatar         *string   `json:"avatar"`
	CreatorUserID  uint      `json:"creator_user_id"`
	OrganizationID uint      `json:"organization_id"`
	Priority       uint      `json:"priority"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

func WorkspaceFromEntities(entities []entity.Workspace) []UserWorkspaceResponse {
	list := make([]UserWorkspaceResponse, 0)

	for _, workspace := range entities {
		resp := UserWorkspaceResponse{
			ID:             workspace.ID,
			Title:          workspace.Title,
			Avatar:         workspace.Avatar,
			CreatorUserID:  workspace.CreatorUserID,
			OrganizationID: workspace.OrganizationID,
			Priority:       workspace.Priority,
			CreateAt:       workspace.CreateAt,
			UpdateAt:       workspace.UpdateAt,
		}
		list = append(list, resp)
	}

	return list
}
