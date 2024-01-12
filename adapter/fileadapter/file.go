package fileadapter

import (
	"taskema/pkg/richerror"
	"taskema/repository/filerepo"
)

type Adapter struct {
	fileRepo filerepo.File
}

func New(
	fileRepo filerepo.File,
) Adapter {
	return Adapter{
		fileRepo: fileRepo,
	}
}

func (a Adapter) DoesFileExist(hash string) (bool, error) {

	_, err := a.fileRepo.GetFile(hash)
	if err != nil {
		cErr := err.(richerror.RichError)
		if cErr.Code() == richerror.CodeNotFound {

			return false, nil
		} else {

			return false, err
		}
	}

	return true, nil
}
