package services

import (
	"io"

	"example.com/facebookclone/initializers"
)


func UploadFile(bucketName, objectName string, file io.Reader) error {
	if initializers.InitMinioClient()
}