package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Pastikan path ini sesuai dengan go.mod Anda
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/router"
)

func main() {
	// Konfigurasi port server
	// Anda bisa menggunakan environment variable atau default ke port tertentu
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika tidak ada env var PORT
	}

	// Inisialisasi router
	appRouter := router.NewRouter()

	log.Printf("Starting TubsComic API server on port %s...\n", port)
	log.Printf("Health check available at http://localhost:%s/health\n", port)

	// Mulai HTTP server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), appRouter)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
