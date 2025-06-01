// internal/router/router.go
package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool" // Tambahkan import ini
)

// NewRouter menginisialisasi dan mengembalikan HTTP router baru.
// Sekarang menerima *pgxpool.Pool sebagai argumen.
func NewRouter(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	// Middleware dasar
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"status": "ok", "message": "TubsComic API is healthy!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Di sini nanti kita akan menambahkan grup rute untuk API v1,
	// dan kita akan pass 'db' ke handler yang memerlukannya.
	// Contoh:
	// comicService := service.NewComicService(db) // Akan kita buat nanti
	// comicHandler := http_handler.NewComicHandler(comicService) // Akan kita buat nanti
	// r.Mount("/api/v1", apiV1Routes(comicHandler)) // apiV1Routes juga perlu diubah untuk pass handler

	return r
}

/*
// Contoh untuk rute API v1 (akan kita buat nanti)
// func apiV1Routes(comicHandler *http_handler.ComicHandler) http.Handler {
// r := chi.NewRouter()
// r.Get("/comics", comicHandler.GetAllComics)
// ... tambahkan rute lain di sini
// return r
// }
*/
