package initializers

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinioClient() {
	var err error
	minioURL := os.Getenv("MINIO_URL")
	// minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	// minioSecretKey := os.Getenv("MINIO_SECRET_KEY")

	minioAccessKey := os.Getenv("MINIO_ROOT_USER")
	minioSecretKey := os.Getenv("MINIO_ROOT_PASSWORD")

	if minioURL == "" || minioAccessKey == "" || minioSecretKey == "" {
		log.Fatalf("MINIO_URL, MINIO_ACCESS_KEY or MINIO_SECRET_KEY environment variable not set.")
	}

	MinioClient, err = minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Failed to create Minio Client: %v", err)
	}
}
