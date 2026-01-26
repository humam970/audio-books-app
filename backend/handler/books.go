package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequestBody[CreateBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	params, err := req.ToParams()
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	book, err := h.repo.CreateBook(r.Context(), params)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, book)
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.repo.GetBook(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get book", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, book)
}

func (h *Handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.ListBooks(r.Context())
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, books)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[UpdateBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	book, err := h.repo.UpdateBook(r.Context(), req.ToParams(id))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, book)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteBook(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
