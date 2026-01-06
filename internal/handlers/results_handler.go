package handlers

import (
	"auto-marksheet-reader/internal/services"
	"auto-marksheet-reader/internal/utils"
	"net/http"
)

type ResultsHandler struct {
	Service *services.MarksheetService
}

func (h *ResultsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	results, err := h.Service.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch results")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, results)
}