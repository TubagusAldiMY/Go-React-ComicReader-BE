// internal/core/service/genre_service.go
package service

import (
	"context"
	"log"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/domain" // Sesuaikan path
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/port"   // Sesuaikan path
)

type genreService struct {
	genreRepo port.GenreRepository // Dependensi ke interface repository
}

// NewGenreService membuat instance baru dari genreService.
func NewGenreService(genreRepo port.GenreRepository) *genreService {
	return &genreService{genreRepo: genreRepo}
}

// ListAll mengambil semua genre melalui repository.
// Di sini bisa ditambahkan logika bisnis lain jika ada.
func (s *genreService) ListAll(ctx context.Context) ([]domain.Genre, error) {
	log.Println("GenreService: Call ListAll") // Contoh logging
	genres, err := s.genreRepo.List(ctx)
	if err != nil {
		// Di sini bisa ada penanganan error spesifik service atau logging tambahan
		log.Printf("GenreService: Error calling genreRepo.List: %v\n", err)
		return nil, err // Kembalikan error sebagaimana adanya atau bungkus dengan error service
	}
	return genres, nil
}
