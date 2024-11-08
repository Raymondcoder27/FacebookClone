package services

import (
	"io"

	"example.com/facebookclone/initializers"
)


func DownloadFile(bucketName, objectName string) ([]byte, error) {
	// Initialize minio client object.
	minioClient := initializers.MinioClient

	//download the object from minio
	object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	//Read the object data into a byte buffer
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, object)
	if err != nil {
		return nil
	}