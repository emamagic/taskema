package useradapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/userrepo"
	"taskema/service/hashingservice"
)

type Adapter struct {
	userRepo   userrepo.User
	hashingSvc hashingservice.Service
}

func New(
	userRepo userrepo.User,
	hashingSvc hashingservice.Service,
) Adapter {
	return Adapter{
		userRepo:   userRepo,
		hashingSvc: hashingSvc,
	}
}

func (a Adapter) DoesUserExistByEmail(email string) (bool, error) {
	op := "sqladapter.DoesUserExistByEmail"

	_, err := a.userRepo.GetByEmail(email)
	if err != nil {
		rErr := err.(richerror.RichError)
		if rErr.Code() == richerror.CodeNotFound {

			return false, nil
		}

		return false,
			richerror.New(op).WithError(err)
	}

	return true, nil
}

func (a Adapter) DoesUserExistByEmailPass(email string, password string) (bool, error) {
	op := "sqladapter.DoesUserExistByEmailPass"

	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		rErr := err.(richerror.RichError)
		if rErr.Code() == richerror.CodeNotFound {

			return false, nil
		}

		return false,
			richerror.New(op).WithError(err)
	}

	if err := a.hashingSvc.HashPassCompare(user.PasswordHash, password); err != nil {
		
		return false, nil
	}

	return true, nil
}
