package services

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func GetFileFromForm(c *gin.Context) (file multipart.File, header *multipart.FileHeader, err error)
