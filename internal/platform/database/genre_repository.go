// internal/platform/database/genre_repository.go
package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/domain" // Sesuaikan path
	"github.com/jackc/pgx/v5/pgxpool"
)

type genreRepository struct {
	db *pgxpool.Pool
}

// NewGenreRepository membuat instance baru dari genreRepository.
func NewGenreRepository(db *pgxpool.Pool) *genreRepository {
	return &genreRepository{db: db}
}

// List mengambil semua genre dari database.
func (r *genreRepository) List(ctx context.Context) ([]domain.Genre, error) {
	query := `SELECT id, name, slug, created_at, updated_at FROM genres ORDER BY name ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying genres: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var genres []domain.Genre
	for rows.Next() {
		var g domain.Genre
		if err := rows.Scan(&g.ID, &g.Name, &g.Slug, &g.CreatedAt, &g.UpdatedAt); err != nil {
			log.Printf("Error scanning genre row: %v\n", err)
			return nil, err
		}
		genres = append(genres, g)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error after iterating genre rows: %v\n", err)
		return nil, err
	}

	if genres == nil { // Pastikan mengembalikan slice kosong, bukan nil, jika tidak ada data
		return []domain.Genre{}, nil
	}

	return genres, nil
}

func (r *genreRepository) Create(ctx context.Context, genre *domain.Genre) error {
	query := `INSERT INTO genres (id, name, slug, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, genre.ID, genre.Name, genre.Slug, genre.CreatedAt, genre.UpdatedAt)
	if err != nil {
		log.Printf("Error creating genre in DB: %v\n", err)
		return err
	}
	return nil
}

// GetBySlug mengambil satu genre berdasarkan slug-nya.
func (r *genreRepository) GetBySlug(ctx context.Context, slug string) (*domain.Genre, error) {
	query := `SELECT id, name, slug, created_at, updated_at FROM genres WHERE slug = $1`
	var g domain.Genre

	err := r.db.QueryRow(ctx, query, slug).Scan(&g.ID, &g.Name, &g.Slug, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("No genre found with slug %s: %v\n", slug, err)
			return nil, domain.ErrDataNotFound // Kita akan definisikan error ini
		}
		log.Printf("Error querying genre by slug %s: %v\n", slug, err)
		return nil, err
	}
	return &g, nil
}
