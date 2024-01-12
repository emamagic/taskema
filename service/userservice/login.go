package userservice

import (
	"taskema/param"
	"taskema/pkg/richerror"
)

func (s Service) Login(req param.UserLoginRequest) (param.UserLoginResponse, error) {
	op := "userservice.Login"

	user, gErr := s.repo.GetByEmail(req.Email)
	if gErr != nil {

		return param.UserLoginResponse{},
			richerror.New(op).WithError(gErr)
	}

	accessToken, aErr := s.authSvc.GenerateAccessToken(user.ID, user.RoleID)
	if aErr != nil {

		return param.UserLoginResponse{},
			richerror.New(op).WithError(aErr)
	}

	refreshToken, rErr := s.authSvc.GenerateRefreshToken(user.ID, user.RoleID)
	if rErr != nil {

		return param.UserLoginResponse{},
			richerror.New(op).WithError(rErr)
	}

	return param.UserLoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
