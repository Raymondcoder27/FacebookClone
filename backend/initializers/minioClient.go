package initializers

import (
	"os"

	"github.com/minio/minio-go/v7"
)

var minioClient *minio.Client

func InitMinioClient() {
	var err error

	minioURL := os.Getenv("MINIO_URL")
}
