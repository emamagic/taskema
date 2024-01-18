package param

import (
	"taskema/entity"
)

type UserOrganizationCreateRequest struct {
	Title         string  `json:"title"`
	Avatar        *string `json:"avatar"`
	CreatorUserID uint
}

type UserOrganizationGetAllRequest struct {
	UserID uint `json:"user_id"`
}

type UserOrganizationDeleteRequest struct {
	OrganizationID uint `json:"organization_id"`
}

type UserOrganizationResponse struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Avatar        *string `json:"avatar"`
	CreatorUserID uint    `json:"creator_user_id"`
	CreateAt      int64   `json:"create_at"`
	UpdateAt      int64   `json:"update_at"`
}

func OrganizationFromEntities(entities []entity.Organization) []UserOrganizationResponse {
	list := make([]UserOrganizationResponse, 0)

	for _, org := range entities {
		resp := UserOrganizationResponse{
			ID:            org.ID,
			Title:         org.Title,
			Avatar:        org.Avatar,
			CreatorUserID: org.CreatorUserID,
			CreateAt:      org.CreateAt.Unix(),
			UpdateAt:      org.UpdateAt.Unix(),
		}
		list = append(list, resp)
	}

	return list
}
