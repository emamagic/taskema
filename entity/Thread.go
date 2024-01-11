package entity

import "time"

type Thread struct {
	ID            uint
	Title         string
	Avatar        *string
	CreatorUserID uint
	MessageIDs    []uint
	UserIds       []uint
	IsMuted       bool
	IsActive      bool
	CreateAt      time.Time
	UpdateAt      time.Time
}
