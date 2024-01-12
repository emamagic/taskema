package entity

import "time"


type Board struct {
	ID            uint
	Title         string
	Avatar        *string
	CreatorUserID uint
	WorkspaceID   uint
	Priority      uint
	CreateAt      time.Time
	UpdateAt      time.Time
}
