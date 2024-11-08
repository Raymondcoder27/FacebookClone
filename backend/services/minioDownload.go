package services


func DownloadFile(bucketName, objectName string) ([]byte, error) {
	// Initialize minio client object.
	client, err := minio.New(
		os.Getenv("MINIO_HOST"),
		os.Getenv("MIN