package entity

import "time"


type Channel struct {
	ID             uint
	Title          string
	Avatar         *string
	userIDS        []uint
	CreatorUserID  uint
	OrganizationID uint
	CreateAt       time.Time
	UpdateAt       time.Time
}
