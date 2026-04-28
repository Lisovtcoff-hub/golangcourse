package dto

import (
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/domain"
)

type UpdateProfileInput struct {
	ID             string  `json:"id"`
	Name           *string `json:"name"`
	Age            *int    `json:"age"`
	Email          *string `json:"email"`
	Phone          *string `json:"phone"`
	IdempotencyKey string  `json:"idempotency_key"`
}

func (i UpdateProfileInput) Validate() error {
	if i.Name == nil && i.Age == nil && i.Email == nil && i.Phone == nil {
		return domain.ErrAllFieldsForUpdateEmpty
	}

	if i.IdempotencyKey == "" {
		return domain.ErrIdempotencyKeyRequired
	}

	return nil
}
