package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "gitlab.golang-school.ru/potok-2/lessons/lesson-24/gen/grpc/profile_v1"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/render"
)

func (h Handlers) CreateProfile(ctx context.Context, i *pb.CreateProfileInput) (*pb.CreateProfileOutput, error) {
	input := dto.CreateProfileInput{
		Name:  i.GetName(),
		Age:   int(i.GetAge()),
		Email: i.GetEmail(),
		Phone: i.GetPhone(),
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateProfileOutput{
		Id: output.ID.String(),
	}, nil
}
