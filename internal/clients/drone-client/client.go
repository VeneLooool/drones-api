package drone_client

import (
	"context"
	"fmt"
	
	"github.com/VeneLooool/drones-api/internal/config"
	"github.com/VeneLooool/drones-api/internal/model"
	"github.com/VeneLooool/drones-api/internal/pb/drone-client/drone-client/api/v1/drones"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	droneClient drones.DronesClient
}

func New(ctx context.Context, cfg *config.DroneClientConfig) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Host, cfg.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient()")
	}

	return &Client{
		droneClient: drones.NewDronesClient(conn),
	}, nil
}

func (c *Client) StartDroneMission(ctx context.Context, drone model.Drone, mission model.Mission) error {
	_, err := c.droneClient.StartDroneMission(ctx, &drones.StartDroneMission_Request{
		Id:      drone.ID,
		Mission: transformMissionToProto(mission),
	})
	if err != nil {
		return errors.Wrap(err, "c.droneClient.StartDroneMission()")
	}
	return nil
}

func transformMissionToProto(mission model.Mission) *drones.Mission {
	return &drones.Mission{
		Coordinates: transformCoordinatesToProto(mission.Coordinates),
	}
}

func transformCoordinatesToProto(coordinates model.Coordinates) []*drones.Coordinate {
	protoCoordinates := make([]*drones.Coordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		protoCoordinates = append(protoCoordinates, &drones.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return protoCoordinates
}
