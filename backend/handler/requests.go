package handler

import (
	"bookserve/repo"
	"time"

	"github.com/google/uuid"
)

type CreateAuthorRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (r *CreateAuthorRequest) ToParams() repo.CreateAuthorParams {
	return repo.CreateAuthorParams{
		Name: r.Name,
		Bio:  r.Bio,
	}
}

type UpdateAuthorRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (r *UpdateAuthorRequest) ToParams(authorID uuid.UUID) repo.UpdateAuthorParams {
	return repo.UpdateAuthorParams{
		Name: r.Name,
		Bio:  r.Bio,
		ID:   authorID,
	}
}

type AddAuthorToBookRequest struct {
	AuthorID uuid.UUID `json:"author_id"`
}

func (r *AddAuthorToBookRequest) ToParams(bookID uuid.UUID) repo.AddAuthorToBookParams {
	return repo.AddAuthorToBookParams{
		AuthorID: r.AuthorID,
		BookID:   bookID,
	}
}

type RemoveAuthorFromBookRequest struct {
	AuthorID uuid.UUID `json:"author_id"`
}

func (r *RemoveAuthorFromBookRequest) ToParams(bookID uuid.UUID) repo.RemoveAuthorFromBookParams {
	return repo.RemoveAuthorFromBookParams{
		AuthorID: r.AuthorID,
		BookID:   bookID,
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateNarratorRequest struct {
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}

func (r *CreateNarratorRequest) ToParams() repo.CreateNarratorParams {
	return repo.CreateNarratorParams{
		Name: r.Name,
		Bio:  r.Bio,
	}
}

type UpdateNarratorRequest struct {
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}

func (r *UpdateNarratorRequest) ToParams(narratorID uuid.UUID) repo.UpdateNarratorParams {
	return repo.UpdateNarratorParams{
		Name: r.Name,
		Bio:  r.Bio,
		ID:   narratorID,
	}
}

type AddNarratorToBookRequest struct {
	NarratorID uuid.UUID `json:"narrator_id"`
}

func (r *AddNarratorToBookRequest) ToParams(bookID uuid.UUID) repo.AddNarratorToBookParams {
	return repo.AddNarratorToBookParams{
		NarratorID: r.NarratorID,
		BookID:     bookID,
	}
}

type RemoveNarratorFromBookRequest struct {
	NarratorID uuid.UUID `json:"narrator_id"`
}

func (r *RemoveNarratorFromBookRequest) ToParams(bookID uuid.UUID) repo.RemoveNarratorFromBookParams {
	return repo.RemoveNarratorFromBookParams{
		NarratorID: r.NarratorID,
		BookID:     bookID,
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateGenreRequest struct {
	Name string `json:"name"`
}

type AddGenreToBookRequest struct {
	GenreID uuid.UUID `json:"genre_id"`
}

func (r *AddGenreToBookRequest) ToParams(bookID uuid.UUID) repo.AddGenreToBookParams {
	return repo.AddGenreToBookParams{
		GenreID: r.GenreID,
		BookID:  bookID,
	}
}

type RemoveGenreFromBookRequest struct {
	GenreID uuid.UUID `json:"genre_id"`
}

func (r *RemoveGenreFromBookRequest) ToParams(bookID uuid.UUID) repo.RemoveGenreFromBookParams {
	return repo.RemoveGenreFromBookParams{
		GenreID: r.GenreID,
		BookID:  bookID,
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateChapterRequest struct {
	Title      string `json:"title"`
	StartTime  int32  `json:"start_time"`
	EndTime    int32  `json:"end_time"`
	OrderIndex int32  `json:"order_index"`
}

func (r *CreateChapterRequest) ToParams(bookID uuid.UUID) repo.CreateChapterParams {
	return repo.CreateChapterParams{
		Title:      r.Title,
		StartTime:  r.StartTime,
		EndTime:    r.EndTime,
		OrderIndex: r.OrderIndex,
		BookID:     bookID,
	}
}

type UpdateChapterRequest struct {
	Title      string `json:"title"`
	StartTime  int32  `json:"start_time"`
	EndTime    int32  `json:"end_time"`
	OrderIndex int32  `json:"order_index"`
}

func (r *UpdateChapterRequest) ToParams(chapterID uuid.UUID) repo.UpdateChapterParams {
	return repo.UpdateChapterParams{
		Title:      r.Title,
		StartTime:  r.StartTime,
		EndTime:    r.EndTime,
		OrderIndex: r.OrderIndex,
		ID:         chapterID,
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateBookRequest struct {
	Title           string
	DurationSeconds int32
	Rating          float64
	ReleaseDate     string
	CoverImageUrl   string
	AudioPreviewUrl string
	IsAbridged      bool
}

func (r *CreateBookRequest) ToParams() (repo.CreateBookParams, error) {
	releaseData, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return repo.CreateBookParams{}, err
	}

	params := repo.CreateBookParams{
		Title:           r.Title,
		DurationSeconds: r.DurationSeconds,
		Rating:          r.Rating,
		ReleaseDate:     releaseData,
		CoverImageUrl:   r.CoverImageUrl,
		AudioPreviewUrl: r.AudioPreviewUrl,
		IsAbridged:      r.IsAbridged,
	}

	return params, nil
}

type UpdateBookRequest struct {
	Title      string
	Rating     float64
	IsAbridged bool
}

func (r *UpdateBookRequest) ToParams(bookID uuid.UUID) repo.UpdateBookParams {
	return repo.UpdateBookParams{
		Title:      r.Title,
		Rating:     r.Rating,
		IsAbridged: r.IsAbridged,
		ID:         bookID,
	}
}
