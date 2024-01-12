package workspaceadapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/workspacerepo"
)

type Adapter struct {
	workspaceRepo workspacerepo.Workspace
}

func New(
	workspaceRepo workspacerepo.Workspace,
) Adapter {
	return Adapter{
		workspaceRepo: workspaceRepo,
	}
}

func (a Adapter) DoesWorkspaceExist(workspaceID uint) (bool, error) {

	_, err := a.workspaceRepo.GetWorkspaceByID(workspaceID)
	if err != nil {
		cErr := err.(richerror.RichError)
		if cErr.Code() == richerror.CodeNotFound {

			return false, nil
		} else {

			return false, err
		}
	}

	return true, nil
}
