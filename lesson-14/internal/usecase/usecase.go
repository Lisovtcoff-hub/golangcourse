package usecase

import (
	"context"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/kafka_produce"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/redis"

	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/domain"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka_produce.Message) error
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	GetProfile(ctx context.Context, id uuid.UUID) (domain.Profile, error)
}

type Profile struct {
	postgres Postgres
	kafka    Kafka
	redis    Redis
}

func NewProfile(postgres *postgres.Pool, kafka *kafka_produce.Producer, redis *redis.Client) *Profile {
	return &Profile{
		postgres: postgres,
		kafka:    kafka,
		redis:    redis,
	}
}
