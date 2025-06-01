// internal/platform/database/genre_repository.go
package database

import (
	"context"
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
