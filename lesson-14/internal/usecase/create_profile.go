package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/dto"

	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/adapter/kafka_produce"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/domain"
)

func (p *Profile) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	var output dto.CreateProfileOutput

	// Проверяем в Redis ключу идемпотентности
	if p.redis.IsExists(ctx, input.Email) {
		return output, domain.ErrAlreadyExists
	}

	// Создаём профиль
	profile := domain.Profile{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      input.Name,
		Age:       input.Age,
		Email:     input.Email,
	}

	// Валидируем
	err := profile.Validate()
	if err != nil {
		return output, fmt.Errorf("validate profile: %w", err)
	}

	// Сохраняем в БД
	err = p.postgres.CreateProfile(ctx, profile)
	if err != nil {
		return output, fmt.Errorf("create profile in postgres: %w", err)
	}

	// Отправляем в Kafka событие создания профиля
	err = p.kafka.Produce(ctx, kafka_produce.Message{})
	if err != nil {
		return output, fmt.Errorf("kafka produce: %w", err)
	}

	return dto.CreateProfileOutput{
		ID: profile.ID,
	}, nil
}
