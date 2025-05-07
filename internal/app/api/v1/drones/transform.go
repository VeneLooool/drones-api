package drones

import (
	"github.com/VeneLooool/drones-api/internal/model"
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
	proto_model "github.com/VeneLooool/drones-api/internal/pb/api/v1/model"
)

var (
	mapDroneStatusToProto = map[model.DroneStatus]proto_model.DroneStatus{
		model.DroneStatusAvailable:   proto_model.DroneStatus_DRONE_STATUS_AVAILABLE,
		model.DroneStatusInMission:   proto_model.DroneStatus_DRONE_STATUS_IN_MISSION,
		model.DroneStatusCharging:    proto_model.DroneStatus_DRONE_STATUS_CHARGING,
		model.DroneStatusMaintenance: proto_model.DroneStatus_DRONE_STATUS_MAINTENANCE,
		model.DroneStatusOffline:     proto_model.DroneStatus_DRONE_STATUS_OFFLINE,
	}
	mapDroneStatusToModel = map[proto_model.DroneStatus]model.DroneStatus{
		proto_model.DroneStatus_DRONE_STATUS_AVAILABLE:   model.DroneStatusAvailable,
		proto_model.DroneStatus_DRONE_STATUS_IN_MISSION:  model.DroneStatusInMission,
		proto_model.DroneStatus_DRONE_STATUS_CHARGING:    model.DroneStatusCharging,
		proto_model.DroneStatus_DRONE_STATUS_MAINTENANCE: model.DroneStatusMaintenance,
		proto_model.DroneStatus_DRONE_STATUS_OFFLINE:     model.DroneStatusOffline,
	}
)

func transformDronesToProto(drones []model.Drone) []*proto_model.Drone {
	result := make([]*proto_model.Drone, 0, len(drones))
	for _, drone := range drones {
		result = append(result, transformDroneToProto(drone))
	}
	return result
}

func transformDroneToProto(drone model.Drone) *proto_model.Drone {
	return &proto_model.Drone{
		Id:     drone.ID,
		Name:   drone.Name,
		Status: mapDroneStatusToProto[drone.Status],
	}
}

func transformDroneCreateReq(req *desc.CreateDrone_Request) model.Drone {
	if req == nil {
		return model.Drone{}
	}
	return model.Drone{Name: req.GetName(), CreatedBy: req.GetCreatedBy()}
}

func transformDroneUpdateReq(req *desc.UpdateDrone_Request) model.Drone {
	if req == nil {
		return model.Drone{}
	}
	return model.Drone{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Status: mapDroneStatusToModel[req.GetStatus()],
	}
}
