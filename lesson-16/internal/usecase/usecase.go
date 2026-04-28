package usecase

import (
	"context"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/domain"
)

//go:generate mockery

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	CreateProperty(ctx context.Context, property domain.Property) error
	GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error)
	UpdateProfile(ctx context.Context, profile domain.Profile) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error
}

type UseCase struct {
	postgres Postgres
	redis    Redis
}

func New(postgres Postgres, redis Redis) *UseCase {
	return &UseCase{
		postgres: postgres,
		redis:    redis,
	}
}
