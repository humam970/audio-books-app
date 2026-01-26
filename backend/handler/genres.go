package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequestBody[CreateGenreRequest](r)
	if err != nil {
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

func (h *Handler) GetGenre(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) ListGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.ListGenres(r.Context())
	if err != nil {
		http.Error(w, "Failed to list genres", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, genres)
}

func (h *Handler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) AddGenreToBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[AddGenreToBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.AddGenreToBook(r.Context(), req.ToParams(id)); err != nil {
		http.Error(w, "Failed to add genre to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) RemoveGenreFromBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[RemoveGenreFromBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.RemoveGenreFromBook(r.Context(), req.ToParams(id)); err != nil {
		http.Error(w, "Failed to add genre to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListGenresForBook(w http.ResponseWriter, r *http.Request) {
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
