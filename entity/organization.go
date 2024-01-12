package entity

import "time"

// TODO - change the creator-user-id to user-id
type Organization struct {
	ID            uint
	Title         string
	Avatar        *string
	CreatorUserID uint
	CreateAt      time.Time
	UpdateAt      time.Time
}
