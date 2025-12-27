package main

import (
	"fmt"
	"log"
	"net/http"

	"auto-marksheet-reader/internal/config"
	"auto-marksheet-reader/internal/handlers"
	"auto-marksheet-reader/internal/services"
	"auto-marksheet-reader/internal/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Connect to Database
	storage.ConnectDB(cfg.MongoURI, cfg.DBName)

	// 3. Setup Router
	r := mux.NewRouter()

	uploadService := &services.UploadService{}
	uploadHandler := &handlers.UploadHandler{Service: uploadService}

	r.HandleFunc("/api/upload", uploadHandler.UploadFile).Methods("POST")
	// Upload File Route

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running 🚀"))
	}).Methods("GET")

	// 4. Setup CORS (Allows frontend to talk to backend)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Vue/React default port
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	// 5. Start Server
	fmt.Printf("Starting server on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
