package routes

import (
	"auto-marksheet-reader/internal/handlers"
	"auto-marksheet-reader/internal/services"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Initialize Services
	uploadService := &services.UploadService{}
	// ocrService := &services.OCRService{} // Uncomment when Tesseract is ready
	marksheetService := &services.MarksheetService{}

	// Initialize Handlers
	uploadHandler := &handlers.UploadHandler{Service: uploadService}
	resultsHandler := &handlers.ResultsHandler{Service: marksheetService}
	exportHandler := &handlers.ExportHandler{}

	// API Routes
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/upload", uploadHandler.UploadFile).Methods("POST")
	api.HandleFunc("/results", resultsHandler.GetAll).Methods("GET")
	api.HandleFunc("/export/excel", exportHandler.ExportExcel).Methods("GET")
	api.HandleFunc("/export/pdf", exportHandler.ExportPDF).Methods("GET")
}
