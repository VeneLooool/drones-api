package drones

import (
	"context"
	
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDronesByAuthor(ctx context.Context, req *desc.GetDronesByAuthor_Request) (*desc.GetDronesByAuthor_Response, error) {
	drones, err := i.droneUC.GetByAuthor(ctx, req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.GetDronesByAuthor_Response{
		Drones: transformDronesToProto(drones),
	}, nil
}
