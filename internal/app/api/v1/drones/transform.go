package drones

import (
	"github.com/VeneLooool/drones-api/internal/model"
	proto_model "github.com/VeneLooool/drones-api/internal/pb/api/v1/model"
)

func transformDroneToProto(drone model.Drone) *proto_model.Drone {
	return &proto_model.Drone{
		Id:     drone.ID,
		Name:   drone.Name,
		Status: drone.Status.String(),
	}
}
