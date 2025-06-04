// internal/core/domain/errors.go
package domain

import "errors"

var (
	ErrDataNotFound     = errors.New("requested data not found")
	ErrValidationFailed = errors.New("validation failed")
	ErrConflictingData  = errors.New("conflicting data")
	// Tambahkan error domain umum lainnya di sini
)
