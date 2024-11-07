package initializers

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitMinioClient() {
	var err error

	minioURL := os.Getenv("MINIO_URL")
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	if minioURL == "" || minioAccessKey = "" || minioSecretKey = "" {
		log.Fatal("MINIO_URL, MINIO_ACCESS_KEY or MINIO_SECRET_KEY environment variable not set.")
	}

	MinioClient, err = minio.New(minioURL, &minioOptions{
		Creds: credentials.NewStaticV4(minioAccessKey, minioSecretKey, "")
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Failed to create Minio Client: %v", err)
	}
}
