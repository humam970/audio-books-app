package logy

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Global    zerolog.Logger
	Authors   zerolog.Logger
	Narrators zerolog.Logger
	Genres    zerolog.Logger
	Chapters  zerolog.Logger
	Books     zerolog.Logger
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	Global = log.With().Str("layer", "global").Logger()
	Authors = log.With().Str("layer", "authors").Logger()
	Narrators = log.With().Str("layer", "narrators").Logger()
	Genres = log.With().Str("layer", "genres").Logger()
	Chapters = log.With().Str("layer", "chapters").Logger()
	Books = log.With().Str("layer", "books").Logger()
}
