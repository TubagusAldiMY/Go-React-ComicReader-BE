package main

import (
	"fmt"
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/config"
	"log"
	"net/http"
	"os"

	// Pastikan path ini sesuai dengan go.mod Anda
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/router"
)

func main() {
	// Konfigurasi port server
	// Muat konfigurasi aplikasi
	cfg := config.LoadConfig() // cfg akan berisi semua konfigurasi kita
	// Anda bisa menggunakan environment variable atau default ke port tertentu
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika tidak ada env var PORT
	}

	// Inisialisasi router
	appRouter := router.NewRouter() // Nanti kita bisa pass cfg ke NewRouter jika diperlukan

	log.Printf("Starting TubsComic API server on port %s...\n", cfg.Port)
	log.Printf("Health check available at http://localhost:%s/health\n", cfg.Port)
	log.Printf("Database URL: %s (sensitive info, hide in production logs)\n", cfg.DatabaseURL) // Hati-hati menampilkan ini di log produksi
	log.Printf("Supabase URL: %s\n", cfg.SupabaseURL)

	// Mulai HTTP server
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), appRouter)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
