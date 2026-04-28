package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "gitlab.golang-school.ru/potok-2/lessons/lesson-24/gen/grpc/profile_v1"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/render"
)

func (h Handlers) DeleteProfile(ctx context.Context, i *pb.DeleteProfileInput) (*emptypb.Empty, error) {
	input := dto.DeleteProfileInput{
		ID: i.GetId(),
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
