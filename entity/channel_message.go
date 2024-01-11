package entity

import "time"

// if text and fileID have content at the same time, the text would be caption for that fileID
type ChannelMessage struct {
	ID            uint
	Text          string
	CreatorUserID uint
	FileID        uint
	ChannelID     uint
	ThreadID      uint
	CreateAt      time.Time
	UpdateAt      time.Time
}
