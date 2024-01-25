package entity

import "time"

type Column struct {
	ID       uint
	title    string
	priority uint
	board_id uint
	CreateAt time.Time
	UpdateAt time.Time
}
