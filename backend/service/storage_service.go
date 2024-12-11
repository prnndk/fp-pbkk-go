package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/prnndk/final-project-golang-pbkk/dto"
)

type (
	StorageService interface {
		HandleImageUpload(ctx context.Context, request dto.UploadFileRequest) (dto.UploadFileResponse, error)
	}

	storageService struct {
	}
)

func NewStorageService() StorageService {
	return &storageService{}
}

func (ss *storageService) HandleImageUpload(ctx context.Context, request dto.UploadFileRequest) (dto.UploadFileResponse, error) {
	// Retrieve the file from the request
	file, err := request.File.Open()
	if err != nil {
		return dto.UploadFileResponse{}, err
	}
	defer file.Close()

	// Validate the file type (example: only allow jpg and png)
	ext := filepath.Ext(request.File.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return dto.UploadFileResponse{}, err
	}

	// Generate a random filename based on timestamp
	randomFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Create a file path within the assets directory
	filePath := filepath.Join("assets", randomFileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return dto.UploadFileResponse{}, err
	}
	defer outFile.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(outFile, file)
	if err != nil {
		return dto.UploadFileResponse{}, err
	}

	// Return the response with the file path
	response := dto.UploadFileResponse{
		Path: filePath,
	}
	return response, nil
}
