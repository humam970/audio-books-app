package genre

import (
	"bookserve/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutesWithHandler(pattern string, r *chi.Mux, h *handler.Handler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.CreateGenre)
		r.Get("/", h.ListGenres)
		r.Get("/{id}", h.GetGenre)
		r.Delete("/{id}", h.DeleteGenre)
	})
}
