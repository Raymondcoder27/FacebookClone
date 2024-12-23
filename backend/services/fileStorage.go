package services

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"

	"example.com/facebookclone/initializers"
	"github.com/minio/minio-go/v7"
)

// func UploadFile(bucketName, objectName string, file io.Reader) error {
// 	if initializers.MinioClient == nil {
// 		return errors.New("minio client is not initialised")
// 	}

// 	//validate bucketname and objectname
// 	if bucketName == "" || objectName == "" {
// 		return errors.New("bucketName and objectName cannot be empty.")
// 	}

// 	//ensure the file is not nil
// 	if file == nil {
// 		return errors.New("file cannot be empty")
// 	}

// 	//read the first 512 bytes to determine the file type
// 	buffer := make([]byte, 512)
// 	_, err := file.Read(buffer)
// 	if err != nil {
// 		return errors.New("unable to read file for content type detection")
// 	}

// 	//detect the content type (e.g image/jpeg, video/mp4)
// 	contentType := http.DetectContentType(buffer)

// 	//reset the file reader position after detecting content type
// 	_, err = file.Seek(0, io.SeekStart)
// 	if err != nil {
// 		return errors.New("unable to reset file reader position")
// 	}

// 	//upload the file to minio
// 	_, err := initializers.MinioClient.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{
// 		ContentType: contentType,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	log.Printf("Successfully uploaded %s to %s/%s\n", objectName, bucketName, objectName)
// 	return nil
// }

func UploadFile(bucketName, objectName string, file io.ReadSeeker) error {
	if initializers.MinioClient == nil {
		return errors.New("minio client is not initialised")
	}

	// Validate bucketName and objectName
	if bucketName == "" || objectName == "" {
		return errors.New("bucketName and objectName cannot be empty.")
	}

	// Ensure the file is not nil
	if file == nil {
		return errors.New("file cannot be empty")
	}

	// Read the first 512 bytes to determine the file type
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return errors.New("unable to read file for content type detection")
	}

	// Detect the content type (e.g., image/jpeg, video/mp4)
	contentType := http.DetectContentType(buffer)

	// Reset the file reader position after detecting content type
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New("unable to reset file reader position")
	}

	// Upload the file to MinIO
	_, err = initializers.MinioClient.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	log.Printf("Successfully uploaded %s to %s/%s\n", objectName, bucketName, objectName)
	return nil
}
