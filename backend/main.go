package main

import (
	"bookserve/handler"
	"bookserve/migrations"
	"bookserve/repo"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
		defer db.Close()

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

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Status is available"))
	})

	mux.HandleFunc("POST   /authors", h.CreateAuthor)
	mux.HandleFunc("GET    /authors", h.ListAuthors)
	mux.HandleFunc("GET    /authors/{id}", h.GetAuthor)
	mux.HandleFunc("PUT    /authors/{id}", h.UpdateAuthor)
	mux.HandleFunc("DELETE /authors/{id}", h.DeleteAuthor)

	mux.HandleFunc("POST   /narrators", h.CreateNarrator)
	mux.HandleFunc("GET    /narrators", h.ListNarrators)
	mux.HandleFunc("GET    /narrators/{id}", h.GetNarrator)
	mux.HandleFunc("PUT    /narrators/{id}", h.UpdateNarrator)
	mux.HandleFunc("DELETE /narrators/{id}", h.DeleteNarrator)

	mux.HandleFunc("POST   /genres", h.CreateGenre)
	mux.HandleFunc("GET    /genres", h.ListGenres)
	mux.HandleFunc("GET    /genres/{id}", h.GetGenre)
	mux.HandleFunc("DELETE /genres/{id}", h.DeleteGenre)

	mux.HandleFunc("POST   /chapters", h.CreateChapter)
	mux.HandleFunc("PUT    /chapters/{id}", h.UpdateChapter)
	mux.HandleFunc("DELETE /chapters/{id}", h.DeleteChapter)

	mux.HandleFunc("POST   /books", h.CreateBook)
	mux.HandleFunc("GET    /books", h.ListBooks)
	mux.HandleFunc("GET    /books/{id}/genres", h.ListGenresForBook)
	mux.HandleFunc("GET    /books/{id}/chapters", h.ListChaptersForBook)
	mux.HandleFunc("GET    /books/{id}", h.GetBook)
	mux.HandleFunc("PUT    /books/{id}", h.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}", h.DeleteBook)

	mux.HandleFunc("POST     /books/{bookID}/authors/{authorID}", h.AddAuthorToBook)
	mux.HandleFunc("DELETE   /books/{bookID}/authors/{authorID}", h.RemoveAuthorFromBook)
	mux.HandleFunc("POST     /books/{bookID}/genres/{genreID}", h.AddGenreToBook)
	mux.HandleFunc("DELETE   /books/{bookID}/genres/{genreID}", h.RemoveGenreFromBook)
	mux.HandleFunc("POST     /books/{bookID}/narrators/{narratorID}", h.AddNarratorToBook)
	mux.HandleFunc("DELETE   /books/{bookID}/narrators/{narratorID}", h.RemoveNarratorFromBook)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Server started on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
