package uservalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (v Validator) Register(req param.UserRegisterRequest) error {
	op := "uservalidation.Register"

	vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, is.Alpha, validation.Length(3, 10)),
		validation.Field(&req.Email, is.EmailFormat),
		validation.Field(&req.Password, validation.Required, validation.Length(3, 10)),
	)

	if vErr != nil {

		return richerror.New(op).WithError(vErr)
	}

	isUserExist, dErr := v.userRepo.DoesUserExistByEmail(req.Email)
	if dErr != nil {

		return richerror.New(op).WithError(dErr)
	}
	if isUserExist {

		return richerror.New(op).
			WithMessage(richerror.MsgErrorUserAlreadyExist).
			WithCode(richerror.CodeInvalid)
	}

	if req.Avatar != nil {

		isAvatarExist, aErr := v.fileRepo.DoesFileExist(*req.Avatar)
		if aErr != nil {
			return richerror.New(op).WithError(aErr)
		}

		if !isAvatarExist {
			return richerror.New(op).
				WithMessage(richerror.MsgErrorAvatarNotValid).
				WithCode(richerror.CodeNotFound)
		}

	}

	return nil
}
