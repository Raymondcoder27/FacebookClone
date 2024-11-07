package initializers

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
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

	
}
