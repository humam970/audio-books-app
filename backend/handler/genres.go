package handler

import (
	"bookserve/repo"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CreateGenreRequest struct {
	Name string `json:"name"`
}

func (h *handler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	req := CreateGenreRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	genre, err := h.repo.CreateGenre(r.Context(), req.Name)
	if err != nil {
		http.Error(w, "Failed to create genre", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, genre)
}

func (h *handler) GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid genre id", http.StatusBadRequest)
		return
	}

	genre, err := h.repo.GetGenre(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Genre not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get genre", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, genre)
}

func (h *handler) ListGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.ListGenres(r.Context())
	if err != nil {
		http.Error(w, "Failed to list genres", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, genres)
}

func (h *handler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid genre id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteGenre(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete genere", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) AddGenreToBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := uuid.Parse(r.PathValue("book_id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	genreID, err := uuid.Parse(r.PathValue("genre_id"))
	if err != nil {
		http.Error(w, "Invalid genre id", http.StatusBadRequest)
		return
	}

	arg := repo.AddGenreToBookParams{
		BookID:  bookID,
		GenreID: genreID,
	}

	if err := h.repo.AddGenreToBook(r.Context(), arg); err != nil {
		http.Error(w, "Failed to add genre to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) RemoveGenreFromBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := uuid.Parse(r.PathValue("book_id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	genreID, err := uuid.Parse(r.PathValue("genre_id"))
	if err != nil {
		http.Error(w, "Invalid genre id", http.StatusBadRequest)
		return
	}

	arg := repo.RemoveGenreFromBookParams{
		BookID:  bookID,
		GenreID: genreID,
	}

	if err := h.repo.RemoveGenreFromBook(r.Context(), arg); err != nil {
		http.Error(w, "Failed to add genre to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) ListGenresForBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	genres, err := h.repo.GetGenresForBook(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get genres for book", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, genres)
}
