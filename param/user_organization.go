package param

import (
	"taskema/entity"
	"time"
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
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Avatar        *string   `json:"avatar"`
	CreatorUserID uint      `json:"creator_user_id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
}

func OrganizationFromEntities(entities []entity.Organization) []UserOrganizationResponse {
	list := make([]UserOrganizationResponse, 0)

	for _, org := range entities {
		resp := UserOrganizationResponse{
			ID:            org.ID,
			Title:         org.Title,
			Avatar:        org.Avatar,
			CreatorUserID: org.CreatorUserID,
			CreateAt:      org.CreateAt,
			UpdateAt:      org.UpdateAt,
		}
		list = append(list, resp)
	}

	return list
}
