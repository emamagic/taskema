package userservice

import (
	"taskema/entity"
	"taskema/service/authservice"
	"taskema/service/hashingservice"
)

type Repository interface {
	Register(entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByID(id uint) (entity.User, error)
}

type Service struct {
	repo        Repository
	hashPassSvc hashingservice.Service
	authSvc     authservice.Service
}

func New(
	repo Repository,
	hashPassSvc hashingservice.Service,
	authSvc authservice.Service,
) Service {
	return Service{
		repo:        repo,
		hashPassSvc: hashPassSvc,
		authSvc:     authSvc,
	}
}
