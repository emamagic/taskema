package entity

import "time"

type File struct {
	ID            int
	Hash          string
	Path          string
	UserCreatorID uint
	CreateAt      time.Time
	UpdateAt      time.Time
}
