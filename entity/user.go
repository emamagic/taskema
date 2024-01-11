package entity

import "time"

type User struct {
	ID           uint
	Name         string
	Avatar	     *string
	Email        string
	PasswordHash string
	RoleID       uint 
	CreateAt     time.Time
	UpdateAt     time.Time
}
