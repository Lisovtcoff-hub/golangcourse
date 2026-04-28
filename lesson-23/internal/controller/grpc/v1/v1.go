package v1

import (
	pb "gitlab.golang-school.ru/potok-2/lessons/lesson-23/gen/grpc/profile_v1"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-23/internal/usecase"
)

type Handlers struct {
	pb.UnimplementedProfileV1Server
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
