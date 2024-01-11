package orgvalidation

type FileRepo interface {
	DoesFileExist(hash string) (bool, error)
}

type Validation struct {
	fileRepo FileRepo
}

func New(fileRepo FileRepo) Validation {
	return Validation{fileRepo: fileRepo}
}
