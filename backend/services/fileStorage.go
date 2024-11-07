package services

import (
	"errors"
	"io"

	"example.com/facebookclone/initializers"
)

func UploadFile(bucketName, objectName string, file io.Reader) error {
	if initializers.MinioClient == nil {
		return errors.New("minio client is not initialised")
	}

	//validate bucketname and objectname
	if bucketName == "" || objectName == "" {
		return errors.New("bucketName and objectName cannot be empty.")
	}
}
