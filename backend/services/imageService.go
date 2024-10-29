package services

import (
	"fmt"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"example.com/facebookclone/models"
	"github.com/disintegration/imaging"
)

// UpdateImage processes the image upload, deletes the old image if it exists, and saves the new one
func UpdateImage(model *models.Post, fileHeader *multipart.FileHeader, width, height, left, top int) (*models.Post, error) {
	// Define the placeholder image path (this can be customized)
	placeholderImagePath := "./public/user-placeholder.png"
	publicPath := "public/"

	// Check if the model already has an image and delete the old one (if it's not the placeholder)
	if model.Image != "" && model.Image != placeholderImagePath {
		currentImagePath := publicPath + filepath.Base(model.Image)
		if _, err := os.Stat(currentImagePath); err == nil {
			if err := os.Remove(currentImagePath); err != nil {
				return nil, fmt.Errorf("could not delete existing image: %v", err)
			}
		}
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open uploaded file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, err := imaging.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("could not decode image: %v", err)
	}

	// Crop the image if dimensions are provided
	if width > 0 && height > 0 {
		img = imaging.Crop(img, image.Rect(left, top, left+width, top+height))
	}

	// Generate a unique name for the image
	extension := filepath.Ext(fileHeader.Filename)
	imageName := fmt.Sprintf("%d%s", time.Now().Unix(), extension)
	savePath := filepath.Join(publicPath, imageName)

	// Save the image
	err = imaging.Save(img, savePath)
	if err != nil {
		return nil, fmt.Errorf("could not save image: %v", err)
	}

	// Update the model's image field
	model.Image = "public/images/" + imageName

	return model, nil
}
