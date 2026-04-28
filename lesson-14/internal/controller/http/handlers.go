package http

import (
	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/usecase"
)

// Обработчики HTTP запросов
type Handlers struct {
	profileService *usecase.Profile
}

func NewHandlers(profileService *usecase.Profile) *Handlers {
	return &Handlers{
		profileService: profileService,
	}
}
