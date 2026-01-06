package main

import (
	"fmt"
	"log"
	"net/http"

	"auto-marksheet-reader/internal/config"
	"auto-marksheet-reader/internal/routes" // Import the new routes
	"auto-marksheet-reader/internal/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect DB
	storage.ConnectDB(cfg.MongoURI, cfg.DBName)

	// 3. Setup Router
	r := mux.NewRouter()

	// Register all routes from the routes package
	routes.RegisterRoutes(r)

	// Health Check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is active"))
	}).Methods("GET")

	// 4. CORS Setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	// 5. Start Server
	fmt.Printf("Starting server on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
