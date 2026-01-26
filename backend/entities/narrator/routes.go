package narrator

import (
	"bookserve/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutesWithHandler(pattern string, r *chi.Mux, h *handler.Handler) {
	if r == nil {
		panic("Cannot register routes to nil mux")
	}

	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.CreateNarrator)
		r.Get("/", h.ListNarrators)
		r.Get("/{id}", h.GetNarrator)
		r.Put("/{id}", h.UpdateNarrator)
		r.Delete("/{id}", h.DeleteNarrator)
	})
}
