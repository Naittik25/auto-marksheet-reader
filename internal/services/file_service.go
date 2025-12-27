package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadService handles file operations
type UploadService struct{}

// SaveFile saves the uploaded file to the local "uploads" directory
func (s *UploadService) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// 1. Create the uploads folder if it doesn't exist
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// 2. Generate a unique filename (timestamp + original name)
	// Example: 17099232_marksheet.jpg
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	filepath := filepath.Join(uploadDir, filename)

	// 3. Create the destination file
	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// 4. Copy the uploaded content to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filepath, nil
}
