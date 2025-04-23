package drones

import (
	"context"

	"github.com/VeneLooool/drones-api/internal/model"
)

type DroneUC interface {
	Create(ctx context.Context, drone model.Drone) (model.Drone, error)
	Update(ctx context.Context, drone model.Drone) (model.Drone, error)
	Get(ctx context.Context, id uint64) (model.Drone, error)
	Delete(ctx context.Context, id uint64) error
}
