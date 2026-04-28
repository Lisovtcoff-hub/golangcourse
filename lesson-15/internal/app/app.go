package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/config"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/adapter/redis"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/controller/http"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/usecase"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/httpserver"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/metrics"
	pgpool "gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/postgres"
	redislib "gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/redis"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/router"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/transaction"
)

func Run(ctx context.Context, c config.Config) error { //nolint:funlen
	// Postgres
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	transaction.Init(pgPool)

	// Redis
	redisClient, err := redislib.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redislib.New: %w", err)
	}

	// UseCase
	uc := usecase.New(
		postgres.New(),
		redis.New(redisClient),
	)

	// Metrics
	httpMetrics := metrics.NewHTTPServer()

	// HTTP
	r := router.New()
	http.ProfileRouter(r, uc, httpMetrics)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("App started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig // wait signal

	log.Info().Msg("App got signal to stop")

	// Controllers close
	httpServer.Close()

	// Adapters close
	redisClient.Close()
	pgPool.Close()

	log.Info().Msg("App stopped!")

	return nil
}
