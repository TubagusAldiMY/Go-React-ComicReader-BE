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

/*
// Kita bisa membuat helper untuk response JSON agar lebih konsisten
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload) // Error handling bisa ditambahkan
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Maka di ListGenres:
// if err != nil {
//    log.Printf("GenreHandler: Error calling genreService.ListAll: %v\n", err)
//    respondWithError(w, http.StatusInternalServerError, "Failed to retrieve genres")
//    return
// }
// respondWithJSON(w, http.StatusOK, genres)
*/
