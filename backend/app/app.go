package app

import (
	"bookserve/entities/author"
	"bookserve/entities/book"
	"bookserve/entities/genre"
	"bookserve/entities/narrator"
	"bookserve/env"
	"bookserve/handler"
	"bookserve/migrations"
	"bookserve/repo"
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
)

type Application struct {
	Pool   *pgxpool.Pool
	Router *chi.Mux
	Logger zerolog.Logger
}

func (a *Application) Migrate() error {
	db, err := sql.Open("pgx", env.Config.GetDSN())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := migrations.MigrateFS(ctx, db, migrations.FS, "."); err != nil {
		return err
	}

	return db.Close()
}

func (a *Application) Mount() error {
	a.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.RealIP)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.Timeout(60 * time.Second))

	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Status is available"))
	})

	queries := repo.New(a.Pool)
	h := handler.New(queries)

	author.RegisterRoutesWithHandler("/authors", a.Router, h)
	narrator.RegisterRoutesWithHandler("/narrators", a.Router, h)
	genre.RegisterRoutesWithHandler("/genres", a.Router, h)
	book.RegisterRoutesWithHandler("/books", a.Router, h)

	return nil
}

func (a *Application) Run() error {
	server := &http.Server{
		Addr:         env.Config.Port,
		Handler:      a.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	a.Logger.Printf("Server Has Started At Addr %s", env.Config.Port)
	return server.ListenAndServe()
}
