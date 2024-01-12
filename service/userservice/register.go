package userservice

import (
	"taskema/entity"
	"taskema/param"
	"taskema/pkg/richerror"
)

func (s Service) Register(req param.UserRegisterRequest) (param.UserRegisterResponse, error) {
	op := "userservice.Register"

	passwordHash, hErr := s.hashPassSvc.HashPassGenerate(req.Password)
	if hErr != nil {
		return param.UserRegisterResponse{},
			richerror.New(op).WithError(hErr)
	}

	user := entity.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Avatar:       req.Avatar,
	}

	createdUser, rErr := s.repo.Register(user)
	if rErr != nil {
		return param.UserRegisterResponse{},
			richerror.New(op).WithError(rErr)
	}

	return param.UserRegisterResponse{
		ID:   createdUser.ID,
		Name: createdUser.Name,
	}, nil
}
