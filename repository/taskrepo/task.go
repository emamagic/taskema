package taskrepo

import "taskema/entity"

type Task interface {
	CreateTask(task entity.Task) (uint, error)
	GetAllTaskByBoardID(boardID uint) ([]entity.Task, error)
	DeleteTaskByID(taskID uint) error
	GetTaskByID(taskID uint) (entity.Task, error)
}
