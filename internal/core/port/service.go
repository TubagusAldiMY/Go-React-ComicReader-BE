// internal/core/port/service.go
package port

import (
	"context"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/domain" // Sesuaikan path
)

// GenreService mendefinisikan operasi logika bisnis untuk genre.
type GenreService interface {
	ListAll(ctx context.Context) ([]domain.Genre, error)
	// Di masa depan: Create, GetByID, Update, Delete
}
