package handler

import (
	"bookserve/repo"
)

type handler struct {
	repo *repo.Queries
}

func New(queries *repo.Queries) *handler {
	return &handler{
		repo: queries,
	}
}
