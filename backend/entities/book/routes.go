package book

import (
	"bookserve/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutesWithHandler(pattern string, r *chi.Mux, h *handler.Handler) {
	if r == nil {
		panic("Cannot register routes to nil mux")
	}

	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.CreateBook)
		r.Post("/{id}/chapters", h.CreateChapter)
		r.Post("/{id}/authors", h.AddAuthorToBook)
		r.Post("/{id}/narrators}", h.AddNarratorToBook)
		r.Post("/{id}/genres", h.AddGenreToBook)

		r.Get("/", h.ListBooks)
		r.Get("/{id}", h.GetBook)
		r.Get("/{id}/authors", h.ListChaptersForBook)
		r.Get("/{id}/narrators", h.ListChaptersForBook)
		r.Get("/{id}/genres", h.ListGenresForBook)
		r.Get("/{id}/chapters", h.ListChaptersForBook)

		r.Put("/{id}", h.UpdateBook)
		r.Put("/{bookID}/chapters/{chapterID}", h.UpdateChapter)

		r.Delete("/{id}", h.DeleteBook)
		r.Delete("/{bookID}/chapters/{chapterID}", h.DeleteChapter)
		r.Delete("/{bookID}/authors/{authorID}", h.RemoveAuthorFromBook)
		r.Delete("/{bookID}/narrators/{narratorID}", h.RemoveNarratorFromBook)
		r.Delete("/{bookID}/genres/{genreID}", h.RemoveGenreFromBook)
	})
}
