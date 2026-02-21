package main

import (
	"bookserve/app"
	"bookserve/env"
	"context"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error

	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.New(context.Background(), env.Config.GetDSN())
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	app := app.Application{
		Pool:   pool,
		Router: chi.NewMux(),
		Logger: log.Logger,
	}

	if err := app.Migrate(); err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to migrate")
	}

	if err := app.Mount(); err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to mount")
	}

	if err := app.Run(); err != nil {
		log.Logger.Fatal().Err(err).Msg("server error")
	}

	log.Logger.Info().Msg("Exiting application")
}
