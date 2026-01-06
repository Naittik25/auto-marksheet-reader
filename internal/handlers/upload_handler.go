package handlers

import (
	"auto-marksheet-reader/internal/services"
	"auto-marksheet-reader/internal/utils"
	"fmt"
	"net/http"
)

type UploadHandler struct {
	Service *services.UploadService
}

func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// 1. Parse and Get File
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid file")
		return
	}
	defer file.Close()

	// 2. Save File
	path, err := h.Service.SaveFile(file, header)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not save file")
		return
	}

	fmt.Println("✅ File saved at:", path)
	fmt.Println("🤖 Starting OCR processing...")

	// 3. Perform OCR
	// ocrService := &services.OCRService{}
	// extractedText, err := ocrService.PerformOCR(path)

	extractedText := "OCR is temporarily disabled while we build Frontend."

	if err != nil {
		fmt.Println("❌ OCR Error:", err)
		extractedText = "Error: " + err.Error()
	} else {
		fmt.Println("✅ Text Extracted!")
	}

	// 4. Return Response (USING extractedText HERE)
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "File processed successfully",
		"path":    path,
		"text":    extractedText, // Go requires this variable to be used here
	})
}
