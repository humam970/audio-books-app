package handler

import (
	"bookserve/repo"
	"database/sql"
	"net/http"
)

type handler struct {
	repo *repo.Queries
}

func New(db *sql.DB) handler {
	return handler{
		repo: repo.New(db),
	}
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	arg := repo.CreateAuthorParams{Name: "", Bio: ""}
	author, err := h.repo.CreateAuthor(r.Context(), arg)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	author, err := h.repo.GetAuthor(r.Context())
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	err := h.repo.ListAuthors()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	err := h.repo.UpdateAuthor()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteAuthor()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) AddAuthorToBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.AddAuthorToBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) RemoveAuthorFromBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.RemoveAuthorFromBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) CreateNarrator(w http.ResponseWriter, r *http.Request) {
	err := h.repo.CreateNarrator()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) GetNarrator(w http.ResponseWriter, r *http.Request) {
	err := h.repo.GetNarrator()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) ListNarrators(w http.ResponseWriter, r *http.Request) {
	narrators, err := h.repo.ListNarrators()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) UpdateNarrator(w http.ResponseWriter, r *http.Request) {
	narrator, err := h.repo.UpdateNarrator()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) DeleteNarrator(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteNarrator()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) AddNarratorToBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.AddNarratorToBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) RemoveNarratorFromBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.RemoveNarratorFromBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := h.repo.CreateGenre()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) GetGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := h.repo.GetGenre()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) ListGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.ListGenres()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteGenre()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) AddGenreToBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.AddGenreToBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) RemoveGenreFromBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.RemoveGenreFromBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) GetGenresForBook(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.GetGenresForBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book, err := h.repo.CreateBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := h.repo.GetBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.ListBooks()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	book, err := h.repo.UpdateBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) CreateChapter(w http.ResponseWriter, r *http.Request) {
	chapter, err := h.repo.CreateChapter()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) ListChaptersByBook(w http.ResponseWriter, r *http.Request) {
	chapters, err := h.repo.ListChaptersByBook()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) UpdateChapter(w http.ResponseWriter, r *http.Request) {
	chapter, err := h.repo.UpdateChapter()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *handler) DeleteChapter(w http.ResponseWriter, r *http.Request) {
	err := h.repo.DeleteChapter()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
