package taskrepo

import "taskema/entity"

type Task interface {
	CreateTask(task entity.Task) (uint, error)
	GetAllTaskByColumnID(columnID uint) ([]entity.Task, error)
	DeleteTaskByID(taskID uint) error
	GetTaskByID(taskID uint) (entity.Task, error)
}
