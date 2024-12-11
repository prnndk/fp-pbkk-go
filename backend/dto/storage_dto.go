package dto

import (
	"errors"
	"mime/multipart"
)

const (
	MESSAGE_FAILED_UPLOAD_FILE  = "Failed uploading file to server"
	MESSAGE_SUCCESS_UPLOAD_FILE = "Success upload file to server"
)

var (
	ErrUploadFile = errors.New("failed to upload file")
)

type (
	UploadFileRequest struct {
		File *multipart.FileHeader `json:"file" form:"file" binding:"required"`
	}

	UploadFileResponse struct {
		Path string `json:"path"`
	}
)
