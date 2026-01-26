package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) CreateChapter(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[CreateChapterRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	chapter, err := h.repo.CreateChapter(r.Context(), req.ToParams(id))
	if err != nil {
		http.Error(w, "Failed to create chapter", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, chapter)
}

func (h *Handler) ListChaptersForBook(w http.ResponseWriter, r *http.Request) {
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

/*
decide if i should use the order index with the book id

	/books/{book_id}/chapters/{chapter_id}

or if i should i use the chapter number instead

	/books/{book_id}/chapters/{chapter_number}
*/
func (h *Handler) UpdateChapter(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[UpdateChapterRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	chapter, err := h.repo.UpdateChapter(r.Context(), req.ToParams(id))
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

func (h *Handler) DeleteChapter(w http.ResponseWriter, r *http.Request) {
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
