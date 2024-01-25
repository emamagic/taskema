package taskservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

type Repository interface {
	CreateTask(task entity.Task) (uint, error)
	GetAllTaskByBoardID(boardID uint) ([]entity.Task, error)
	DeleteTaskByID(taskID uint) error
}

type Service struct {
	repo Repository
}

func New(
	repo Repository,
) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateTask(req param.UserTaskCreateRequest) (uint, error) {
	op := "taskservice.CreateTask"

	task := entity.Task{
		Title:          req.Title,
		Avatar:         req.Avatar,
		CreatorUserID:  req.CreatorUserID,
		Description:    req.Description,
		ColumnID:        req.ColumnID,
		AssignedUserID: req.AssignedUserID,
		DueDate:        &req.DueDate,
	}

	id, err := s.repo.CreateTask(task)

	if err != nil {

		return 0, richerror.New(op).WithError(err)
	}

	return id, nil
}

func (s Service) GetAllTaskByBoardID(req param.UserTaskGetAllRequest) ([]param.UserTaskResponse, error) {
	op := "taskservice.GetAllTaskByBoardID"

	tasks, err := s.repo.GetAllTaskByBoardID(req.UserID)
	if err != nil {

		return nil,
			richerror.New(op).WithError(err)
	}
	return param.TaskFromEntities(tasks), nil
}

func (s Service) DeleteTaskByID(req param.UserTaskDeleteRequest) error {
	op := "taskservice.DeleteTaskByID"

	if err := s.repo.DeleteTaskByID(req.TaskID); err != nil {

		return richerror.New(op).WithError(err)
	}

	return nil
}
