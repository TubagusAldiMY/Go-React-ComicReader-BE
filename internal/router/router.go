// internal/router/router.go
package router

import (
	"encoding/json"
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/service"
	http_handler "github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/handler/http"
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/platform/database"
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

	// ... (inisialisasi middleware, health check, dan layers genre tetap sama) ...
	genreRepo := database.NewGenreRepository(db)
	genreService := service.NewGenreService(genreRepo)
	genreHandler := http_handler.NewGenreHandler(genreService)

	r.Route("/api/v1", func(r chi.Router) {
		// --- Rute untuk Admin ---
		r.Route("/admin", func(r chi.Router) {
			// (Di sini nanti kita bisa tambahkan middleware autentikasi admin)
			r.Post("/genres", genreHandler.CreateGenre)
			r.Put("/genres/{genreSlug}", genreHandler.UpdateGenre)
		})

		// --- Rute Publik ---
		r.Get("/genres", genreHandler.ListGenres)
		r.Get("/genres/{genreSlug}", genreHandler.GetGenreBySlug)
	})

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
