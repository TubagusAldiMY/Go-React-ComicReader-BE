// internal/core/domain/genre.go
package domain

import (
	"time"

	"github.com/google/uuid" // Jika Anda menggunakan UUID
)

// Genre merepresentasikan entitas genre komik.
type Genre struct {
	ID        uuid.UUID `json:"id"` // Menggunakan github.com/google/uuid
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
