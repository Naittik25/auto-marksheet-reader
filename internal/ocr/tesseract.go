package ocr

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// TesseractClient handles the OCR operations using the command line
type TesseractClient struct{}

// NewTesseractClient creates a new client
func NewTesseractClient() *TesseractClient {
	return &TesseractClient{}
}

// ExtractText calls the "tesseract" command directly from your installed software
func (t *TesseractClient) ExtractText(imagePath string) (string, error) {
	// We run the command: tesseract <image> stdout
	// "stdout" tells Tesseract to print text to the screen instead of a file
	cmd := exec.Command("C:\\Program Files\\Tesseract-OCR\\tesseract.exe", imagePath, "stdout")

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("tesseract command failed: %v. Make sure Tesseract is installed and in your PATH", err)
	}

	// Clean up the text (remove extra whitespace)
	text := strings.TrimSpace(out.String())
	return text, nil
}

// Close is empty because we don't keep a connection open in this method
func (t *TesseractClient) Close() {
	// No cleanup needed for command line method
}
