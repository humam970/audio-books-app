package main

import (
	"context"
	"database/sql"

	"bookserve/entities/author"
	"bookserve/entities/book"
	"bookserve/entities/genre"
	"bookserve/entities/narrator"
	l "bookserve/logy"
	"net/http"
	"os"
	"os/signal"
	"time"

	"bookserve/handler"
	"bookserve/migrations"
	"bookserve/repo"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connStr := "postgres://test:test@localhost:5433/test?sslmode=disable"

	{
		db, err := sql.Open("pgx", connStr)
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = db.Close()
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := migrations.MigrateFS(ctx, db, migrations.FS, "."); err != nil {
			panic(err)
		}
	}

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	queries := repo.New(pool)
	h := handler.New(queries)

	r := chi.NewMux()

	//nolint
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	// r.Use(l.Middleware)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Status is available"))
	})

	author.RegisterRoutesWithHandler("/authors", r, h)
	narrator.RegisterRoutesWithHandler("/narrators", r, h)
	genre.RegisterRoutesWithHandler("/genres", r, h)
	book.RegisterRoutesWithHandler("/books", r, h)

	// nolint
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Base.Printf("Server started on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Base.Err(err).Msg("listen: %s\n")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	l.Base.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		l.Base.Fatal().Msgf("Server forced to shutdown: %v", err)
	}

	l.Base.Println("Server exiting")
}
