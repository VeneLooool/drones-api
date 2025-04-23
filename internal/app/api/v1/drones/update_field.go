package drones

import (
	"context"
	
	"github.com/VeneLooool/drones-api/internal/model"
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateDrone(ctx context.Context, req *desc.UpdateDrone_Request) (*desc.UpdateDrone_Response, error) {
	field, err := i.droneUC.Update(ctx, model.Drone{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Status: model.DroneStatus(req.GetStatus()),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UpdateDrone_Response{
		Drone: transformDroneToProto(field),
	}, nil
}
