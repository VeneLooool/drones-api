package drones

import (
	"context"
	"errors"
	
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"github.com/VeneLooool/drones-api/internal/pkg/error_hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDroneByID(ctx context.Context, req *desc.GetDroneByID_Request) (*desc.GetDroneByID_Response, error) {
	drone, err := i.droneUC.Get(ctx, req.GetId())
	if err != nil {
		code := codes.Internal
		if errors.Is(err, error_hub.ErrDroneNotFound) {
			code = codes.NotFound
		}
		return nil, status.Error(code, err.Error())
	}
	return &desc.GetDroneByID_Response{
		Drone: transformDroneToProto(drone),
	}, nil
}
