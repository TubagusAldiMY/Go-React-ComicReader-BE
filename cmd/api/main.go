// cmd/api/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/config"
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/platform/database"
	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	dbPool, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("FATAL: Could not connect to database: %v", err)
	}
	defer database.CloseDB() // database.DB akan di-close di sini jika masih menggunakan global var di CloseDB

	// Sekarang dbPool digunakan saat memanggil NewRouter
	appRouter := router.NewRouter(dbPool)

	log.Printf("Starting TubsComic API server on port %s...\n", cfg.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: appRouter,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  120 * time.Second,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
		log.Println("Server exited properly")
	}()

	log.Printf("Server listening on %s. Press Ctrl+C to quit.", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	}
}
