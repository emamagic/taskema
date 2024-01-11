package hashingservice

import (
	"taskema/pkg/richerror"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) HashPassGenerate(password string) (string, error) {
	op := "hashingservice.HashPassGenerate"

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	return string(hashPassword), nil
}

func (s Service) HashPassCompare(hashPassword string, password string) error {
	op := "hashingservice.HashPassCompare"

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return richerror.New(op).
			WithMessage(err.Error()).
			WithCode(richerror.CodeUnexpected)
	}
	return nil
}

func (s Service) GenerateRandomUniqueID() string {

	id := uuid.New()
	idString := id.String()

	return idString
}
