package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/config" // Sesuaikan path ke config Anda
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB adalah variabel global untuk menyimpan connection pool (atau bisa di-pass sebagai dependency)
// Untuk sekarang, kita buat sebagai variabel yang bisa diakses,
// tapi idealnya di-inject ke mana pun dibutuhkan.
var DB *pgxpool.Pool

// ConnectDB menginisialisasi koneksi ke database PostgreSQL menggunakan connection pool.
func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {
	log.Println("Attempting to connect to database...")

	// Konfigurasi parsing dari DatabaseURL
	dbConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database configuration: %w", err)
	}

	// Anda bisa mengatur parameter pool di sini jika perlu
	// dbConfig.MaxConns = 10 // Contoh
	// dbConfig.MinConns = 2  // Contoh
	// dbConfig.MaxConnLifetime = time.Hour // Contoh
	// dbConfig.MaxConnIdleTime = 30 * time.Minute // Contoh

	// Membuat connection pool
	// Menggunakan context.Background() karena ini adalah setup awal.
	// Timeout bisa ditambahkan jika diperlukan.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Timeout 10 detik untuk koneksi
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Ping database untuk memastikan koneksi berhasil.
	// Menggunakan context.Background() untuk ping, atau context dengan timeout pendek.
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second) // Timeout 5 detik untuk ping
	defer pingCancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close() // Tutup pool jika ping gagal
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database!")
	DB = pool // Assign ke variabel global (jika menggunakan pendekatan ini)
	return pool, nil
}

// CloseDB menutup connection pool database.
// Fungsi ini bisa dipanggil saat aplikasi shutdown.
func CloseDB() {
	if DB != nil {
		log.Println("Closing database connection pool...")
		DB.Close()
	}
}
