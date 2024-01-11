package uservalidation

type UserRepo interface {
	DoesUserExistByEmail(email string) (bool, error)
	DoesUserExistByEmailPass(email string, password string) (bool, error)
}

type FileRepo interface {
	DoesFileExist(hash string) (bool, error)
}

type Validator struct {
	userRepo UserRepo
	fileRepo FileRepo
}

func New(
	userRepo UserRepo,
	fileRepo FileRepo,
) Validator {
	return Validator{
		userRepo: userRepo,
		fileRepo: fileRepo,
	}
}
