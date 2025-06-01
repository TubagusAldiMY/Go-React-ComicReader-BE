package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Config menyimpan semua konfigurasi aplikasi.
// Nilai-nilai ini dibaca dari environment variables.
type Config struct {
	Port        string
	DatabaseURL string // Kita akan gabungkan info DB menjadi satu URL

	SupabaseURL string
	SupabaseKey string

	// Tambahkan field lain jika diperlukan
}

// LoadConfig memuat konfigurasi dari environment variables.
// Jika file .env ada, ia akan dimuat terlebih dahulu (berguna untuk development lokal).
func LoadConfig() *Config {

	// Log current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: Could not get current working directory: %v", err)
	} else {
		log.Printf("Current working directory: %s", cwd)
		envPath := filepath.Join(cwd, ".env") // Path yang diharapkan untuk .env
		log.Printf("Attempting to load .env file from: %s", envPath)
	}

	// Coba muat file .env jika ada
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v. Relying on OS environment variables.", err)
	} else {
		log.Println(".env file loaded successfully.")
	}
	// Coba muat file .env jika ada (biasanya untuk local development)
	// Dalam production, variabel environment biasanya diatur langsung di server/platform.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env, relying on OS environment variables")
	}

	// Dapatkan Database URL dari komponen atau langsung
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		dbHost := getEnv("DB_HOST", "localhost")
		dbPort := getEnv("DB_PORT", "5432")
		dbUser := getEnv("DB_USER", "postgres")         // Ganti default jika perlu
		dbPassword := getEnv("DB_PASSWORD", "password") // Ganti default jika perlu
		dbName := getEnv("DB_NAME", "tubcomic_db")      // Ganti default jika perlu
		dbSSLMode := getEnv("DB_SSLMODE", "disable")
		driver := getEnv("DB_DRIVER", "postgres")

		// Format: postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
		databaseURL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			driver, dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: databaseURL,
		SupabaseURL: getEnv("SUPABASE_URL", ""),
		SupabaseKey: getEnv("SUPABASE_KEY", ""),
	}
}

// getEnv membaca environment variable atau mengembalikan nilai default.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if fallback == "" && (key == "SUPABASE_URL" || key == "SUPABASE_KEY") { // Contoh: Wajib ada
		log.Fatalf("FATAL: Environment variable %s not set and no fallback provided.", key)
	}
	return fallback
}

// getEnvAsInt membaca environment variable sebagai integer atau mengembalikan nilai default.
func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("FATAL: Environment variable %s is not a valid integer: %s", key, valueStr)
		return fallback // tidak akan tercapai karena Fatalf
	}
	return value
}
