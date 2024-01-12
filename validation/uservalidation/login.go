package uservalidation

import (
	"taskema/param"
	"taskema/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (v Validator) Login(req param.UserLoginRequest) error {
	op := "uservalidation.Login"

	vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Email, is.EmailFormat),
		validation.Field(&req.Password, validation.Required, validation.Length(3, 10)),
	)

	if vErr != nil {

		return richerror.New(op).WithError(vErr)
	}

	isExist, dErr := v.userRepo.DoesUserExistByEmailPass(req.Email, req.Password)
	if dErr != nil {

		return richerror.New(op).WithError(dErr)
	}
	if !isExist {
		return richerror.New(op).
			WithMessage(richerror.MsgErrorNoSuchUserExist).
			WithCode(richerror.CodeInvalid)
	}

	return nil
}
