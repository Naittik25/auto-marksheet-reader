package handlers

import (
	"auto-marksheet-reader/internal/utils"
	"net/http"
)

// Placeholder for Export Logic
type ExportHandler struct{}

func (h *ExportHandler) ExportExcel(w http.ResponseWriter, r *http.Request) {
	// TODO: Integrate Excel Service here
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Excel download started"})
}

func (h *ExportHandler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	// TODO: Integrate PDF Service here
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "PDF download started"})
}
