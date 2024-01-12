package filerepo

import "taskema/entity"

type File interface {
	GetFile(hash string) (entity.File, error)
	StoreFile(file entity.File) (string, error)
}
