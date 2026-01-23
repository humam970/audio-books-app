package main

import (
	"context"
	"database/sql"

	"bookserve/entities/author"
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
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connStr := ""

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
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(l.Middleware)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Status is available"))
	})

	// r.Route("/authors", func(r chi.Router) {
	// 	r.Post("/", h.CreateAuthor)
	// 	r.Get("/", h.ListAuthors)
	// 	r.Get("/{id}", h.GetAuthor)
	// 	r.Put("/{id}", h.UpdateAuthor)
	// 	r.Delete("/{id}", h.DeleteAuthor)
	// })
	author.RegisterRoutes(r)

	r.Route("/narrators", func(r chi.Router) {
		r.Post("/", h.CreateNarrator)
		r.Get("/", h.ListNarrators)
		r.Get("/{id}", h.GetNarrator)
		r.Put("/{id}", h.UpdateNarrator)
		r.Delete("/{id}", h.DeleteNarrator)
	})

	r.Route("/genres", func(r chi.Router) {
		r.Post("/", h.CreateGenre)
		r.Get("/", h.ListGenres)
		r.Get("/{id}", h.GetGenre)
		r.Delete("/{id}", h.DeleteGenre)
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", h.CreateBook)
		r.Post("/{id}/chapters", h.CreateChapter)
		r.Post("/{bookID}/authors/{authorID}", h.AddAuthorToBook)
		r.Post("/{bookID}/genres/{genreID}", h.AddGenreToBook)
		r.Post("/{bookID}/narrators/{narratorID}", h.AddNarratorToBook)

		r.Get("/", h.ListBooks)
		r.Get("/{id}", h.GetBook)
		r.Get("/{id}/genres", h.ListGenresForBook)
		r.Get("/{id}/chapters", h.ListChaptersForBook)

		r.Put("/{id}", h.UpdateBook)
		r.Put("/{bookID}/chapters/{chapterID}", h.UpdateChapter)

		r.Delete("/{id}", h.DeleteBook)
		r.Delete("/{bookID}/chapters/{chapterID}", h.DeleteChapter)
		r.Delete("/{bookID}/authors/{authorID}", h.RemoveAuthorFromBook)
		r.Delete("/{bookID}/genres/{genreID}", h.RemoveGenreFromBook)
		r.Delete("/{bookID}/narrators/{narratorID}", h.RemoveNarratorFromBook)
	})

	//nolint:exhaustruct
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Global.Printf("Server started on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Global.Err(err).Msg("listen: %s\n")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	l.Global.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		l.Global.Fatal().Msgf("Server forced to shutdown: %v", err)
	}

	l.Global.Println("Server exiting")
}

// export const bookKeys = {
// 	all: ["books"] as const,

// 	// List and Search
// 	lists: () => [...bookKeys.all, "list"] as const,
// 	search: (query: string) => [...bookKeys.all, "search", query] as const,

// 	// Specific Book Scopes
// 	details: () => [...bookKeys.all, "detail"] as const,
// 	detail: (id: string) => [...bookKeys.details(), id] as const,

// 	// Sub-resource: Chapters
// 	chapters: (bookId: string) =>
// 		[...bookKeys.detail(bookId), "chapters"] as const,
// 	chapter: (bookId: string, chapterId: string) =>
// 		[...bookKeys.chapters(bookId), chapterId] as const,

// 	// Sub-resource: Metadata (Authors, Genres, Narrators)
// 	// We group these under "metadata" so you can invalidate all book relationships at once if needed
// 	authors: (bookId: string) => [...bookKeys.detail(bookId), "authors"] as const,
// 	genres: (bookId: string) => [...bookKeys.detail(bookId), "genres"] as const,
// 	narrators: (bookId: string) =>
// 		[...bookKeys.detail(bookId), "narrators"] as const,
// };
