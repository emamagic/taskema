package userrepo

import "taskema/entity"

type User interface {
	GetByID(id uint) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Register(user entity.User) (entity.User, error)
}
