package services

import (
	"auto-marksheet-reader/internal/ocr"
	"fmt"
)

type OCRService struct {
	Tesseract *ocr.TesseractClient
}

// PerformOCR takes a file path, reads it, and returns text
func (s *OCRService) PerformOCR(filePath string) (string, error) {
	// 1. Initialize Tesseract Client
	client := ocr.NewTesseractClient()
	defer client.Close()

	// 2. Extract Text
	text, err := client.ExtractText(filePath)
	if err != nil {
		return "", err
	}

	if text == "" {
		return "", fmt.Errorf("no text found in image")
	}

	return text, nil
}
