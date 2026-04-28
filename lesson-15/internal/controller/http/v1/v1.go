package v1

import "gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/usecase"

type Handlers struct {
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
