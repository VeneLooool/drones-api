package drones

import (
	"context"

	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateDrone(ctx context.Context, req *desc.CreateDrone_Request) (*desc.CreateDrone_Response, error) {
	drone, err := i.droneUC.Create(ctx, transformDroneCreateReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateDrone_Response{
		Drone: transformDroneToProto(drone),
	}, nil
}
