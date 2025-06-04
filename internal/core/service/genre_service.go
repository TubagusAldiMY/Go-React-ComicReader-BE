// internal/core/service/genre_service.go
package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"

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

// generateSlug membuat slug yang SEO-friendly dari string.
func generateSlug(s string) string {
	// Ganti spasi dengan strip, kecilkan semua huruf
	slug := strings.ToLower(s)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Anda bisa menambahkan logika pembersihan karakter non-alfanumerik di sini jika perlu
	// Contoh sederhana:
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func (s *genreService) CreateNewGenre(ctx context.Context, name string) (*domain.Genre, error) {
	log.Printf("GenreService: Call CreateNewGenre with name: %s\n", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("genre name cannot be empty") // Contoh validasi sederhana
	}

	slug := generateSlug(name)
	// Cek duplikasi slug/name bisa ditambahkan di sini dengan memanggil repository

	newGenre := &domain.Genre{
		ID:        uuid.New(), // Generate UUID baru
		Name:      name,
		Slug:      slug,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.genreRepo.Create(ctx, newGenre)
	if err != nil {
		log.Printf("GenreService: Error calling genreRepo.Create: %v\n", err)
		return nil, err
	}

	return newGenre, nil
}

func (s *genreService) FindGenreBySlug(ctx context.Context, slug string) (*domain.Genre, error) {
	log.Printf("GenreService: Call FindGenreBySlug with slug: %s\n", slug)
	genre, err := s.genreRepo.GetBySlug(ctx, slug)
	if err != nil {
		// Tidak perlu log lagi di sini jika repository sudah log
		// Cukup teruskan errornya, atau bungkus jika perlu konteks tambahan
		return nil, err // Ini akan meneruskan domain.ErrDataNotFound jika itu errornya
	}
	return genre, nil
}
