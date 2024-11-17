package service

import (
	"errors"
	"fmt"
	"minio-learn/model"
	"minio-learn/repository"
)

type UploadService interface {
	UploadFile(file model.UploadFile) error
	GetFile(fileName string) (string, error)
}

type uploadService struct {
	repo repository.UploadFileMinio
}

func NewUploadService(repo repository.UploadFileMinio) UploadService {
	return &uploadService{repo}
}

func (u *uploadService) UploadFile(file model.UploadFile) error {

	buffer, err := file.File.Open()
	if err != nil {
		return err
	}

	defer buffer.Close()

	objectName := file.FileName
	fileBuffer := buffer
	contentType := file.File.Header.Get("Content-Type")
	fileSize := file.File.Size

	allowedTYPES := map[string]string{
		"application/pdf": "pdf",
		"image/jpg":       "jpg",
		"image/jpeg":      "jpeg",
		"image/png":       "png",
	}

	isValidType := false
	for _, allowedType := range allowedTYPES {
		if contentType == allowedType {
			isValidType = true
			break
		}
	}

	ext, isValidType := allowedTYPES[contentType]
	if !isValidType {
		return errors.New("file type not allowed")
	}

	fileName := fmt.Sprintf("01/nama-file/%s.%s", objectName, ext)

	if !isValidType {
		return errors.New("file type not allowed")
	}

	return u.repo.UploadFileMinio(fileName, fileBuffer, contentType, fileSize)
}

func (u *uploadService) GetFile(fileName string) (string, error) {

	fileUrl, err := u.repo.GetFileMinio(fileName)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}
