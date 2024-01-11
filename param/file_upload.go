package param

import "mime/multipart"

type FileUploadRequest struct {
	Files         []*multipart.FileHeader
	UserCreatorID uint
}

type FileUploadResponse struct {
	Hash string `json:"hash"`
}
