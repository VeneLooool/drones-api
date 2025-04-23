package drones

import (
	"context"

	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteDrone(ctx context.Context, req *desc.DeleteDrone_Request) (*emptypb.Empty, error) {
	if err := i.droneUC.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
