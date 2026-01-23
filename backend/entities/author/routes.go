package author

import (
	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/authors", func(r chi.Router) {
		// r.Post("/", h.CreateAuthor)
		// r.Get("/", h.ListAuthors)
		// r.Get("/{id}", h.GetAuthor)
		// r.Put("/{id}", h.UpdateAuthor)
		// r.Delete("/{id}", h.DeleteAuthor)
	})
}
