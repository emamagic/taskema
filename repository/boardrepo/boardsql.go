package boardrepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type boardSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) Board {
	return &boardSqlRepo{db: db}
}


func (repo *boardSqlRepo) CreateBoard(board entity.Board) (uint, error) {
	op := "boardsql.CreateBoard"

	result, err := repo.db.Conn().Exec("INSERT INTO boards (title, avatar, creator_user_id, workspace_id) VALUES (?, ?, ?, ?)",
		board.Title, board.Avatar, board.CreatorUserID, board.WorkspaceID)
	if err != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	id, iErr := result.LastInsertId()
	if iErr != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return uint(id), nil
}

func (repo *boardSqlRepo) GetAllBoardByWorkspaceID(workspaceID uint) ([]entity.Board, error) {
	op := "boardsql.GetAllBoardByWorkspaceID"

	rows, qErr := repo.db.Conn().Query("SELECT * FROM boards WHERE workspace_id = ?", workspaceID)
	if qErr != nil {

		return nil,
			richerror.New(op).
				WithMessage(qErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	defer rows.Close()

	list := make([]entity.Board, 0)
	for rows.Next() {
		var board entity.Board
		err := rows.Scan(
			&board.ID,
			&board.Title,
			&board.Avatar,
			&board.CreatorUserID,
			&board.WorkspaceID,
			&board.Priority,
			&board.CreateAt,
			&board.UpdateAt,
		)
		if err != nil {

			return nil,
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeUnexpected)
		}
		list = append(list, board)
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



func (repo *boardSqlRepo) DeleteBoardByID(boardID uint) error {
	op := "boardsql.DeleteBoardByID"

	result, err := repo.db.Conn().Exec("DELETE FROM boards WHERE id = ?", boardID)
	if err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	count, aErr := result.RowsAffected()
	if aErr != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if count == 0 {

		return richerror.New(op).WithMessage("no such board exist").WithCode(richerror.CodeNotFound)
	}

	return nil
}

func (repo *boardSqlRepo) GetBoardByID(boardID uint) (entity.Board, error) {
	op := "boardsql.GetBoardByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM boards WHERE id = ?", boardID)

	var board entity.Board
	if err := row.Scan(
		&board.ID,
		&board.Title,
		&board.Avatar,
		&board.CreatorUserID,
		&board.WorkspaceID,
		&board.Priority,
		&board.CreateAt,
		&board.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.Board{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}
		return entity.Board{},
			richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return board, nil
}
