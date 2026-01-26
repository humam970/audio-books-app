package handler

import (
	"bookserve/repo"
)

type Handler struct {
	repo *repo.Queries
}

func New(queries *repo.Queries) *Handler {
	if queries == nil {
		panic("Cannot create new handler with nil queries")
	}

	return &Handler{
		repo: queries,
	}
}
