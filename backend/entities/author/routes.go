package author

import (
	"bookserve/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutesWithHandler(pattern string, r *chi.Mux, h *handler.Handler) {
	if r == nil {
		panic("Cannot register routes to nil mux")
	}

	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.CreateAuthor)
		r.Get("/", h.ListAuthors)
		r.Get("/{id}", h.GetAuthor)
		r.Put("/{id}", h.UpdateAuthor)
		r.Delete("/{id}", h.DeleteAuthor)
	})
}
