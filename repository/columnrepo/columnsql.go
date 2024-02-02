package columnrepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type columnSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) Column {
	return &columnSqlRepo{db: db}
}

func (repo *columnSqlRepo) CreateColumn(column entity.Column) (uint, error) {
	op := "columnsql.CreateColumn"

	result, err := repo.db.Conn().Exec("INSERT INTO columns (title, creator_user_id, workspace_id) VALUES (?, ?, ?)",
		column.Title, column.CreatorUserID, column.WorkspaceID)
	if err != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	id, iErr := result.LastInsertId()
	if iErr != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return uint(id), nil
}

func (repo *columnSqlRepo) GetAllColumnByWorkspaceID(workspaceID uint) ([]entity.Column, error) {
	op := "columnsql.GetAllColumnsByWorkspaceID"

	rows, qErr := repo.db.Conn().Query("SELECT * FROM columns WHERE workspace_id = ?", workspaceID)
	if qErr != nil {

		return nil,
			richerror.New(op).
				WithMessage(qErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	defer rows.Close()

	list := make([]entity.Column, 0)
	for rows.Next() {
		var column entity.Column
		err := rows.Scan(
			&column.ID,
			&column.Title,
			&column.CreatorUserID,
			&column.Priority,
			&column.WorkspaceID,
			&column.CreateAt,
			&column.UpdateAt,
		)
		if err != nil {

			return nil,
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeUnexpected)
		}
		list = append(list, column)
	}

	if rows.Err() != nil {

		return nil,
			richerror.New(op).
				WithMessage(rows.Err().Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "error happend while scanning one of the row of columns"})
	}

	return list, nil
}

func (repo *columnSqlRepo) DeleteColumnByID(columnID uint) error {
	op := "columnsql.DeleteColumnByID"

	result, err := repo.db.Conn().Exec("DELETE FROM columns WHERE id = ?", columnID)
	if err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	count, aErr := result.RowsAffected()
	if aErr != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if count == 0 {

		return richerror.New(op).WithMessage("no such column exist").WithCode(richerror.CodeNotFound)
	}

	return nil
}

func (repo *columnSqlRepo) GetColumnByID(columnID uint) (entity.Column, error) {
	op := "columnsql.GetColumnByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM columns WHERE id = ?", columnID)

	var column entity.Column
	if err := row.Scan(
		&column.ID,
		&column.Title,
		&column.CreatorUserID,
		&column.Priority,
		&column.WorkspaceID,
		&column.CreateAt,
		&column.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.Column{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}
		return entity.Column{},
			richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return column, nil
}
