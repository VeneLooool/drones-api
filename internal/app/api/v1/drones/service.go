package drones

import (
	desc "github.com/VeneLooool/drones-api/internal/pb/api/v1/drones"
)

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedDronesServer

	droneUC DroneUC
}

// NewService return new instance of Implementation.
func NewService(droneUC DroneUC) *Implementation {
	return &Implementation{
		droneUC: droneUC,
	}
}
