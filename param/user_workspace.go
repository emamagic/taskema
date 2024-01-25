package param

import (
	"taskema/entity"
)

type UserWorkspaceCreateRequest struct {
	CreatorUserID  uint
	Title          string  `json:"title"`
	Avatar         *string `json:"avatar"`
	OrganizationID uint    `json:"organization_id"`
}

type UserWorkspaceGetAllRequest struct {
	UserID         uint
	OrganizationID uint `json:"organization_id"`
}

type UserWorkspaceDeleteRequest struct {
	UserID      uint
	WorkspaceID uint `json:"workspace_id"`
}

type UserWorkspaceResponse struct {
	ID             uint    `json:"id"`
	Title          string  `json:"title"`
	Avatar         *string `json:"avatar"`
	CreatorUserID  uint    `json:"creator_user_id"`
	OrganizationID uint    `json:"organization_id"`
	Priority       uint    `json:"priority"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
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
			CreateAt:       workspace.CreateAt.Unix(),
			UpdateAt:       workspace.UpdateAt.Unix(),
		}
		list = append(list, resp)
	}

	return list
}
