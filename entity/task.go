package entity

import "time"

type Task struct {
	ID             uint
	Title          string
	Avatar         *string
	CreatorUserID  uint
	Description    string
	DueDate        *int64
	BoardID        uint
	AssignedUserID uint
	Priority       uint
	CreateAt       time.Time
	UpdateAt       time.Time
}
