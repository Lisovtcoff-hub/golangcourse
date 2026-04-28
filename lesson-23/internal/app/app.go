package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/config"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/adapter/kafka_producer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/adapter/redis"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/controller/grpc"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/controller/http"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/controller/kafka_consumer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/controller/worker"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/usecase"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/httpserver"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/metrics"
	pgpool "gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/postgres"
	redislib "gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/redis"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/router"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/pkg/transaction"
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

	// Kafka producer
	kafkaProducer := kafka_producer.New(c.KafkaProducer, metrics.NewProcess())

	// UseCase
	uc := usecase.New(
		postgres.New(),
		redis.New(redisClient),
		kafkaProducer,
	)

	// Kafka consumer
	kafkaConsumer := kafka_consumer.New(c.KafkaConsumer, uc)

	// Outbox Kafka worker
	outboxKafkaWorker := worker.NewOutboxKafka(uc, c.OutboxKafka)

	// Metrics
	httpMetrics := metrics.NewHTTPServer()

	// GRPC
	grpcServer, err := grpc.New(c.GRPC, uc)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	// HTTP
	r := router.New()
	http.ProfileRouter(r, uc, httpMetrics)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig // wait signal

	log.Info().Msg("app: got signal to stop")

	// Controllers close
	httpServer.Close()
	grpcServer.Close()
	outboxKafkaWorker.Close()
	kafkaConsumer.Close()

	// Adapters close
	redisClient.Close()
	kafkaProducer.Close()
	pgPool.Close()

	log.Info().Msg("app: stopped")

	return nil
}
