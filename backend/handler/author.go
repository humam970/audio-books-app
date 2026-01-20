package handler

import (
	"bookserve/repo"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	arg := repo.CreateAuthorParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	author, err := h.repo.CreateAuthor(r.Context(), arg)
	if err != nil {
		http.Error(w, "Fialed to add author to db", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, author)
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid author id", http.StatusBadRequest)
		return
	}

	author, err := h.repo.GetAuthor(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Aurhor not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get author", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, author)
}

func (h *handler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.repo.ListAuthors(r.Context())
	if err != nil {
		http.Error(w, "Failed to get authors", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, authors)
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid author id", http.StatusBadRequest)
		return
	}
	_ = id

	arg := repo.UpdateAuthorParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	author, err := h.repo.UpdateAuthor(r.Context(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, author)
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid author id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteAuthor(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) AddAuthorToBook(w http.ResponseWriter, r *http.Request) {
	arg := repo.AddAuthorToBookParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := h.repo.AddAuthorToBook(r.Context(), arg)
	if err != nil {
		http.Error(w, "Failed to add author to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) RemoveAuthorFromBook(w http.ResponseWriter, r *http.Request) {
	arg := repo.RemoveAuthorFromBookParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.RemoveAuthorFromBook(r.Context(), arg); err != nil {
		http.Error(w, "Failed to remove author from book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
