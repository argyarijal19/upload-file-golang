package repository

import (
	"context"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
)

type UploadFileMinio interface {
	UploadFileMinio(fileName string, buffer multipart.File, contentType string, fileSize int64) error
	GetFileMinio(fileName string) (string, error)
}

type uploadFile struct {
	upload *minio.Client
}

func NewUploadRepository(upload *minio.Client) UploadFileMinio {
	return &uploadFile{upload}
}

func (u *uploadFile) UploadFileMinio(fileName string, buffer multipart.File, contentType string, fileSize int64) error {
	var bucket = os.Getenv("BUCKET_NAME")

	_, err := u.upload.PutObject(context.Background(), bucket, fileName, buffer, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *uploadFile) GetFileMinio(fileName string) (string, error) {
	var bucket = os.Getenv("BUCKET_NAME")

	fileUrl, err := u.upload.PresignedGetObject(context.Background(), bucket, fileName, 3*time.Minute, nil)
	if err != nil {
		log.Println("Error presigned URL: ", err)
		return "", err
	}

	return fileUrl.String(), nil
}
