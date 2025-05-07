package drones

import (
	"context"
	"errors"
	
	"github.com/VeneLooool/drones-api/internal/model"
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	proto_model "github.com/VeneLooool/drones-api/internal/pb/api/v1/model"
	"github.com/VeneLooool/drones-api/internal/pkg/error_hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) StartDroneMission(ctx context.Context, req *desc.StartDroneMission_Request) (*emptypb.Empty, error) {
	err := i.droneUC.StartDroneMission(ctx, req.GetId(), transformMissionToModel(req.GetMission()))
	if err != nil {
		code := codes.Internal
		if errors.Is(err, error_hub.ErrDroneNotFound) {
			code = codes.NotFound
		}
		return nil, status.Error(code, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func transformMissionToModel(mission *proto_model.Mission) model.Mission {
	return model.Mission{
		Coordinates: transformCoordinatesToModel(mission.GetCoordinates()),
	}
}

func transformCoordinatesToModel(protoCoordinates []*proto_model.Coordinate) model.Coordinates {
	coordinates := make([]model.Coordinate, 0, len(protoCoordinates))
	for _, coordinate := range protoCoordinates {
		coordinates = append(coordinates, model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return coordinates
}
