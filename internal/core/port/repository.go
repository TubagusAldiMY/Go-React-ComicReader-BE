// internal/core/port/repository.go
package port

import (
	"context"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/domain" // Sesuaikan path
)

// GenreRepository mendefinisikan operasi yang bisa dilakukan pada data genre.
type GenreRepository interface {
	// Di masa depan, kita bisa menambahkan method lain seperti Create, GetByID, Update, Delete
	List(ctx context.Context) ([]domain.Genre, error)
}
