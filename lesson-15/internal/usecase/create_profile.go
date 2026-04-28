package usecase

import (
	"context"
	"fmt"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/transaction"
)

func (u *UseCase) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreateProfile")
	defer span.End()

	var output dto.CreateProfileOutput

	profile, err := domain.NewProfile(input.Name, input.Age, input.Email, input.Phone)
	if err != nil {
		return output, fmt.Errorf("domain.NewProfile: %w", err)
	}

	property := domain.NewProperty(profile.ID, []string{"home", "primary"})

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = u.postgres.CreateProfile(ctx, profile)
		if err != nil {
			return fmt.Errorf("postgres.CreateProfile: %w", err)
		}

		err = u.postgres.CreateProperty(ctx, property)
		if err != nil {
			return fmt.Errorf("postgres.CreateProperty: %w", err)
		}

		return nil
	})
	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.CreateProfileOutput{
		ID: profile.ID,
	}, nil
}
