package drones

import (
	"context"

	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateDrone(ctx context.Context, req *desc.UpdateDrone_Request) (*desc.UpdateDrone_Response, error) {
	field, err := i.droneUC.Update(ctx, transformDroneUpdateReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UpdateDrone_Response{
		Drone: transformDroneToProto(field),
	}, nil
}
