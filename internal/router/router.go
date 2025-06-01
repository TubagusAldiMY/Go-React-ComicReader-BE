package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter menginisialisasi dan mengembalikan HTTP router baru.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware dasar
	r.Use(middleware.RequestID) // Menyuntikkan Request ID unik ke setiap request
	r.Use(middleware.RealIP)    // Menggunakan IP asli dari header X-Forwarded-For atau X-Real-IP
	r.Use(middleware.Logger)    // Log request (method, path, duration, status)
	r.Use(middleware.Recoverer) // Memulihkan dari panic tanpa menghentikan server

	// Atur timeout untuk menghindari request yang menggantung terlalu lama
	// r.Use(middleware.Timeout(60 * time.Second)) // uncomment jika diperlukan

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Siapkan response
		response := map[string]string{"status": "ok", "message": "TubsComic API is healthy!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Di sini nanti kita akan menambahkan grup rute untuk API v1, misalnya /api/v1
	// r.Mount("/api/v1", apiV1Routes())

	return r
}

/*
// Contoh untuk rute API v1 (akan kita buat nanti)
func apiV1Routes() http.Handler {
	r := chi.NewRouter()
	// r.Get("/comics", comicHandler.GetAllComics)
	// ... tambahkan rute lain di sini
	return r
}
*/
