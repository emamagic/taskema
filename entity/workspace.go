package entity

import "time"

type Workspace struct {
	ID             uint
	Title          string
	Avatar         *string
	CreatorUserID  uint
	OrganizationID uint
	Priority       uint
	CreateAt       time.Time
	UpdateAt       time.Time
}
