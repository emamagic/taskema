package taskrepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type taskSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) Task {
	return &taskSqlRepo{db: db}
}

// TODO - user-assigned not implemeted yet
func (repo *taskSqlRepo) CreateTask(task entity.Task) (uint, error) {
	op := "tasksql.CreateTask"

	result, err := repo.db.Conn().Exec("INSERT INTO tasks (title, avatar, creator_user_id, description, board_id, assigned_user_id, due_date) VALUES (?, ?, ?, ?, ?, ?, ?)",
		task.Title, task.Avatar, task.CreatorUserID, task.Description, task.ColumnID, task.AssignedUserID, task.DueDate)
	if err != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	id, iErr := result.LastInsertId()
	if iErr != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return uint(id), nil
}

func (repo *taskSqlRepo) GetAllTaskByBoardID(boardID uint) ([]entity.Task, error) {
	op := "tasksql.GetAllTaskByBoardID"

	rows, qErr := repo.db.Conn().Query("SELECT * FROM tasks WHERE board_id = ?", boardID)
	if qErr != nil {

		return nil,
			richerror.New(op).
				WithMessage(qErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	defer rows.Close()

	list := make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Avatar,
			&task.CreatorUserID,
			&task.Description,
			&task.DueDate,
			&task.ColumnID,
			&task.AssignedUserID,
			&task.Priority,
			&task.CreateAt,
			&task.UpdateAt,
		)
		if err != nil {

			return nil,
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeUnexpected)
		}
		list = append(list, task)
	}

	if rows.Err() != nil {

		return nil,
			richerror.New(op).
				WithMessage(rows.Err().Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "error happend while scanning one of the row of boards"})
	}

	return list, nil
}

func (repo *taskSqlRepo) DeleteTaskByID(taskID uint) error {
	op := "tasksql.DeleteTaskByID"

	result, err := repo.db.Conn().Exec("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	count, aErr := result.RowsAffected()
	if aErr != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if count == 0 {

		return richerror.New(op).WithMessage("no such task exist").WithCode(richerror.CodeNotFound)
	}

	return nil
}

func (repo *taskSqlRepo) GetTaskByID(taskID uint) (entity.Task, error) {
	op := "tasksql.DeleteTaskByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM tasks WHERE id = ?", taskID)

	var task entity.Task
	if err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Avatar,
		&task.CreatorUserID,
		&task.Description,
		&task.DueDate,
		&task.ColumnID,
		&task.AssignedUserID,
		&task.Priority,
		&task.CreateAt,
		&task.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.Task{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}
		return entity.Task{},
			richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return task, nil
}
