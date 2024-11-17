package lib

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnMinio() (*minio.Client, error) {
	endpoint := "localhost:9000"
	accessKeyID := os.Getenv("ACCESS_KEY_MINIO")
	secretAccessKey := os.Getenv("SECRET_KEY_MINIO")
	client, err := minio.New((endpoint), &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
