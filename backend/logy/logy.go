package logy

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	Base      zerolog.Logger
	Authors   zerolog.Logger
	Narrators zerolog.Logger
	Genres    zerolog.Logger
	Chapters  zerolog.Logger
	Books     zerolog.Logger
)

func init() {

	if true {
		//nolint
		Base = zerolog.New(zerolog.ConsoleWriter{
			Out: os.Stderr,
			FormatTimestamp: func(i any) string {
				s := fmt.Sprintf("%v", i)
				t, err := time.Parse(time.RFC3339, s)
				if err != nil {
					return colorizeString(green, s)
				}

				return colorizeString(green, t.Format(time.Kitchen))
			},
			FormatMessage: func(i any) string {
				return colorizeString(blue, fmt.Sprintf("| %s |", i))
			},
		}).With().Timestamp().Logger()
	} else {
		Base = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	Authors = Base.With().Str("layer", "authors").Logger()
	Narrators = Base.With().Str("layer", "narrators").Logger()
	Genres = Base.With().Str("layer", "genres").Logger()
	Chapters = Base.With().Str("layer", "chapters").Logger()
	Books = Base.With().Str("layer", "books").Logger()
}
