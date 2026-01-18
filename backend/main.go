package main

import (
	"bookserve/handler"
	"bookserve/migrations"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	if err := migrations.MigrateFS(db, migrations.FS, "."); err != nil {
		panic(err)
	}

	h := handler.New(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Status is available"))
	})

	mux.HandleFunc("POST   /authors", h.CreateAuthor)
	mux.HandleFunc("GET    /authors", h.GetAuthor)
	mux.HandleFunc("PUT    /authors", h.UpdateAuthor)
	mux.HandleFunc("DELETE /authors", h.DeleteAuthor)

	mux.HandleFunc("POST   /narrators", h.CreateNarrator)
	mux.HandleFunc("GET    /narrators", h.GetNarrator)
	mux.HandleFunc("PUT    /narrators", h.UpdateNarrator)
	mux.HandleFunc("DELETE /narrators", h.CreateAuthor)

	mux.HandleFunc("POST   /books", h.CreateAuthor)
	mux.HandleFunc("GET    /books", h.CreateAuthor)
	mux.HandleFunc("PUT    /books", h.CreateAuthor)
	mux.HandleFunc("DELETE /books", h.CreateAuthor)

	mux.HandleFunc("POST   /chapters", h.CreateAuthor)
	mux.HandleFunc("GET    /chapters", h.CreateAuthor)
	mux.HandleFunc("PUT    /chapters", h.CreateAuthor)
	mux.HandleFunc("DELETE /chapters", h.CreateAuthor)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server Statrted On Port: 8080")
	log.Fatal(server.ListenAndServe())
}
