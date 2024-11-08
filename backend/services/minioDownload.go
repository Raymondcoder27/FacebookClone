package services

import "example.com/facebookclone/initializers"


func DownloadFile(bucketName, objectName string) ([]byte, error) {
	// Initialize minio client object.
	minioClient := initializers.MinioClient

	//download the object from minio
	object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()