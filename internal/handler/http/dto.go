// internal/handler/http/dto.go
package http_handler // atau package yang sesuai dengan handler Anda

// CreateGenreRequest adalah DTO untuk request pembuatan genre.
type CreateGenreRequest struct {
	Name string `json:"name" validate:"required"` // Kita akan tambahkan validasi nanti
}
