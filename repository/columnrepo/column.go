package columnrepo

import "taskema/entity"

type Column interface {
	CreateColumn(column entity.Column) (uint, error)
	GetAllColumnByWorkspaceID(workspaceID uint) ([]entity.Column, error)
	DeleteColumnByID(columnID uint) error
	GetColumnByID(columnID uint) (entity.Column, error)
}
