package filerepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type fileSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) File {
	return &fileSqlRepo{db: db}
}

func (repo *fileSqlRepo) GetFile(hash string) (entity.File, error) {
	op := "filesql.GetFile"

	row := repo.db.Conn().QueryRow("SELECT * FROM files WHERE hash = ?", hash)

	var file entity.File
	if err := row.Scan(&file.ID, &file.Hash, &file.Path, &file.UserCreatorID, &file.CreateAt, &file.UpdateAt); err != nil {
		if err == sql.ErrNoRows {

			return entity.File{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}

		return entity.File{},
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)

	}

	return file, nil
}

func (repo *fileSqlRepo) StoreFile(file entity.File) (string, error) {
	op := "filesql.StoreFile"

	_, err := repo.db.Conn().Exec("INSERT INTO files (hash, path, user_creator_id) values (?, ?, ?)",
		file.Hash, file.Path, file.UserCreatorID)

	if err != nil {

		return "", richerror.New(op).
			WithMessage(err.Error()).
			WithCode(richerror.CodeUnexpected)
	}

	return file.Hash, nil
}
