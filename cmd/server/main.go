package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Auto Marksheet Reader Backend is running on port 8080...")

	// Simple health check route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend is Active!")
	})

	// Start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
