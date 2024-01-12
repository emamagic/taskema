package fileservice

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"taskema/entity"
	"taskema/param"
	"taskema/service/hashingservice"
	"time"
)

type Repository interface {
	GetFile(hash string) (entity.File, error)
	StoreFile(file entity.File) (string, error)
}

type Service struct {
	hashingSvc hashingservice.Service
	repo       Repository
}

func New(
	hashSvc hashingservice.Service,
	repo Repository,
) Service {
	return Service{
		hashingSvc: hashSvc,
		repo:       repo,
	}
}

func (s Service) StoreFile(req param.FileUploadRequest) (param.FileUploadResponse, error) {
	files := req.Files

	// TODO - setup for multiple file uploaded
	// TODO - decrease file size
	for _, file := range files {

		src, err := file.Open()
		if err != nil {
			return param.FileUploadResponse{}, err
		}
		defer src.Close()
		fileNameAndMimeType := strings.Split(file.Filename, ".")
		fileName := fmt.Sprintf("%s_%s.%s", fileNameAndMimeType[0], strconv.FormatInt(time.Now().UnixNano(), 10), fileNameAndMimeType[1])
		uploadPath := filepath.Join("uploads", fileName)
		dst, err := os.Create("uploads/" + fileName)
		if err != nil {
			return param.FileUploadResponse{}, err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return param.FileUploadResponse{}, err
		}

		randomUniqueID := s.hashingSvc.GenerateRandomUniqueID()
		if err != nil {
			return param.FileUploadResponse{}, err
		}

		file := entity.File{
			Hash:          randomUniqueID,
			Path:          uploadPath,
			UserCreatorID: req.UserCreatorID,
		}

		hash, sErr := s.repo.StoreFile(file)
		if sErr != nil {
			return param.FileUploadResponse{}, sErr
		}

		return param.FileUploadResponse{
			Hash: hash,
		}, nil
	}

	// unreachable code
	return param.FileUploadResponse{
		Hash: "",
	}, nil
}

func (s Service) GetFile(hash string) (string, error) {
	file, err := s.repo.GetFile(hash)
	if err != nil {
		return "", err
	}

	return file.Path, nil
}
