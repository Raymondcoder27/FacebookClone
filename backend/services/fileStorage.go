package services

import (
	"context"
	"errors"
	"io"

	"example.com/facebookclone/initializers"
	"github.com/minio/minio-go/v7"
)

func UploadFile(bucketName, objectName string, file io.Reader) error {
	if initializers.MinioClient == nil {
		return errors.New("minio client is not initialised")
	}

	//validate bucketname and objectname
	if bucketName == "" || objectName == "" {
		return errors.New("bucketName and objectName cannot be empty.")
	}

	//upload the file to minio
	_, err := initializers.MinioClient.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{
		ContentType: "image",
	})
}
