package taskadapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/taskrepo"
)

type Adapter struct {
	taskRepo taskrepo.Task
}

func New(
	taskRepo taskrepo.Task,
) Adapter {
	return Adapter{
		taskRepo: taskRepo,
	}
}

func (a Adapter) DoesTaskExist(taskID uint) (bool, error) {

	_, err := a.taskRepo.GetTaskByID(taskID)
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
