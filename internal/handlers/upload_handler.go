package handlers

import (
	"net/http"

	// UPDATE THIS IMPORT to match your go.mod name (either "auto-marksheet-reader" or "github.com/...")
	"auto-marksheet-reader/internal/services"
	"auto-marksheet-reader/internal/utils"
)

type UploadHandler struct {
	Service *services.UploadService
}

// UploadFile handles the POST /upload request
func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// 1. Limit upload size to 10MB
	r.ParseMultipartForm(10 << 20)

	// 2. Get the file from the request
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid file upload")
		return
	}
	defer file.Close()

	// 3. Save the file using the service
	path, err := h.Service.SaveFile(file, header)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not save file")
		return
	}

	// 4. Respond with success
	response := map[string]string{
		"message": "File uploaded successfully",
		"path":    path,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
