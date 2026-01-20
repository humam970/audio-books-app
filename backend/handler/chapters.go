package handler

import (
	"bookserve/repo"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *handler) CreateChapter(w http.ResponseWriter, r *http.Request) {
	arg := repo.CreateChapterParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	chapter, err := h.repo.CreateChapter(r.Context(), arg)
	if err != nil {
		http.Error(w, "Failed to create chapter", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, chapter)
}

func (h *handler) ListChaptersForBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	chapters, err := h.repo.ListChaptersByBook(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get chapters", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, chapters)
}

func (h *handler) UpdateChapter(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid chapter id", http.StatusBadRequest)
		return
	}
	_ = id

	arg := repo.UpdateChapterParams{}
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	chapter, err := h.repo.UpdateChapter(r.Context(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Chapter not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update chapter", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, chapter)
}

func (h *handler) DeleteChapter(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid chapter id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteChapter(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete chapter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
