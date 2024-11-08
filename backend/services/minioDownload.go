package services

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"example.com/facebookclone/initializers"
	// "github.com/minio/minio-go"
	"github.com/minio/minio-go/v7"
)

func DownloadFile(bucketName, objectName string) ([]byte, error) {
	// Initialize minio client object.
	minioClient := initializers.MinioClient

	//download the object from minio
	object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	//Read the object data into a byte buffer
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, object)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DeleteFile(bucketName, objectName string) error {
	// Initialize minio client object.
	minioClient := initializers.MinioClient

	//delete the object from minio
	err := minioClient.RemoveObject(bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Successfully deleted", objectName)

	return nil
}
