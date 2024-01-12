package workspacerepo

import "taskema/entity"

type Workspace interface {
	CreateWorkspace(workspace entity.Workspace) (uint, error)
	GetAllWorkspaceByOrganizationID(organizationID uint, userID uint) ([]entity.Workspace, error)
	DeleteWorkspaceByID(workspaceID uint) error
	GetWorkspaceByID(workspaceID uint) (entity.Workspace, error)
}
