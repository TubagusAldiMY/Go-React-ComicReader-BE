// internal/handler/http/genre_handler.go
package http_handler // Menggunakan http_handler untuk menghindari konflik nama dengan package http standar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/port" // Sesuaikan path
)

type GenreHandler struct {
	genreService port.GenreService // Dependensi ke interface service
}

// NewGenreHandler membuat instance baru dari GenreHandler.
func NewGenreHandler(genreService port.GenreService) *GenreHandler {
	return &GenreHandler{genreService: genreService}
}

// ListGenres menangani request untuk mendapatkan semua genre.
func (h *GenreHandler) ListGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.genreService.ListAll(r.Context())
	if err != nil {
		log.Printf("GenreHandler: Error calling genreService.ListAll: %v\n", err)
		// Kirim response error yang lebih baik di sini nanti
		http.Error(w, "Failed to retrieve genres", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(genres); err != nil {
		log.Printf("GenreHandler: Error encoding genres to JSON: %v\n", err)
		// Jika terjadi error encoding, client mungkin sudah menerima status 200
		// jadi kita tidak bisa mengubah header lagi. Cukup log errornya.
	}
}

// CreateGenre menangani request untuk membuat genre baru.
func (h *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var req CreateGenreRequest // DTO yang kita buat tadi
	// Decode JSON request body ke struct CreateGenreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("GenreHandler: Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validasi sederhana (bisa menggunakan library validator nanti)
	if req.Name == "" {
		http.Error(w, "Genre name is required", http.StatusBadRequest)
		return
	}

	createdGenre, err := h.genreService.CreateNewGenre(r.Context(), req.Name)
	if err != nil {
		log.Printf("GenreHandler: Error calling genreService.CreateNewGenre: %v\n", err)
		// Cek tipe error spesifik dari service jika ada
		if err.Error() == "genre name cannot be empty" { // Contoh penanganan error spesifik
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to create genre", http.StatusInternalServerError)
		}
		return
	}

	// Kirim response sukses dengan data genre yang baru dibuat
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Status 201 Created
	if err := json.NewEncoder(w).Encode(createdGenre); err != nil {
		log.Printf("GenreHandler: Error encoding created genre to JSON: %v\n", err)
	}
}
