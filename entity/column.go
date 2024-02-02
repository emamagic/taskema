package entity

import "time"

type Column struct {
	ID            uint
	Title         string
	CreatorUserID uint
	Priority      uint
	WorkspaceID   uint
	CreateAt      time.Time
	UpdateAt      time.Time
}
