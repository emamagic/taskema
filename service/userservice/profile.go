package userservice

import (
	"taskema/param"
	"taskema/pkg/richerror"
)

func (s Service) GetUserProfile(req param.UserProfileRequest) (param.UserProfileResponse, error) {
	op := "userservice.GetUserProfile"

	user, err := s.repo.GetByID(req.UserID)
	if err != nil {

		return param.UserProfileResponse{},
			richerror.New(op).WithError(err)
	}

	return param.UserProfileResponse{
		ID:     user.ID,
		Name:   user.Name,
		Avatar: user.Avatar,
		Email:  user.Email,
	}, nil
}
