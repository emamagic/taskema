package userrepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type userSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) User {
	return &userSqlRepo{db: db}
}

func (repo *userSqlRepo) GetByID(id uint) (entity.User, error) {
	op := "usersql.GetByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM users WHERE id = ?", id)

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Avatar,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
		&user.CreateAt,
		&user.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.User{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}

		return entity.User{},
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	return user, nil
}

func (repo *userSqlRepo) GetByEmail(email string) (entity.User, error) {
	op := "usersql.GetByEmail"

	row := repo.db.Conn().QueryRow("SELECT * FROM users WHERE email=?", email)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Avatar,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
		&user.CreateAt,
		&user.UpdateAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {

			return entity.User{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}

		return entity.User{},
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	return user, nil
}

func (repo *userSqlRepo) Register(user entity.User) (entity.User, error) {
	op := "usersql.Register"

	result, eErr := repo.db.Conn().Exec("INSERT INTO users(name, avatar, email, password) VALUES(?, ?, ?, ?)",
		user.Name, user.Avatar, user.Email, user.PasswordHash)
	if eErr != nil {

		return entity.User{},
			richerror.New(op).
				WithMessage(eErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	userID, iErr := result.LastInsertId()
	if iErr != nil {

		return entity.User{},
			richerror.New(op).
				WithMessage(iErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	user.ID = uint(userID)

	// returning user_request_param_info of user without create_at, update_at, ...
	return user, nil
}

