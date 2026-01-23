package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

func (h *handler) CreateNarrator(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequestBody[CreateNarratorRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	narrator, err := h.repo.CreateNarrator(r.Context(), req.ToParams())
	if err != nil {
		http.Error(w, "Failed to create narrator", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusCreated, narrator)
}

func (h *handler) GetNarrator(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid narrator id", http.StatusBadRequest)
		return
	}

	narrator, err := h.repo.GetNarrator(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Narrator not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get narrator", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, narrator)
}

func (h *handler) ListNarrators(w http.ResponseWriter, r *http.Request) {
	narrators, err := h.repo.ListNarrators(r.Context())
	if err != nil {
		http.Error(w, "Failed to get narrators", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, narrators)
}

func (h *handler) UpdateNarrator(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid narrator id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[UpdateNarratorRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	narrator, err := h.repo.UpdateNarrator(r.Context(), req.ToParams(id))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Narrator not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update narrator", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, narrator)
}

func (h *handler) DeleteNarrator(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid narrator id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteNarrator(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete narrator", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) AddNarratorToBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[AddNarratorToBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.AddNarratorToBook(r.Context(), req.ToParams(id)); err != nil {
		http.Error(w, "Failed to add author to book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) RemoveNarratorFromBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	req, err := decodeRequestBody[RemoveNarratorFromBookRequest](r)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.RemoveNarratorFromBook(r.Context(), req.ToParams(id)); err != nil {
		http.Error(w, "Failed to remove author from book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
