package entity

import "time"


// just one of the workspaceID and taskID has the value at the same time [ Host (task, workspace) ]
type Conversation struct {
	ID            uint
	Title         string
	Avatar        *string
	CreatorUserID uint
	TaskID        uint
	WorkspaceID   uint
	UserIds       []uint
	IsMuted       bool
	IsActive      bool
	CreateAt      time.Time
	UpdateAt      time.Time
}

