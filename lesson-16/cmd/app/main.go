package main

import (
	"context"

	"github.com/rs/zerolog/log"

	_ "go.uber.org/automaxprocs"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/config"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/app"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/logger"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/otel"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	ctx := context.Background()

	err = otel.Init(ctx, c.OTEL)
	if err != nil {
		log.Fatal().Err(err).Msg("otel.Init")
	}
	defer otel.Close()

	err = app.Run(ctx, c)
	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
}
