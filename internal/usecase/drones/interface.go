package drones

import (
	"context"

	"github.com/VeneLooool/drones-api/internal/model"
)

type Repo interface {
	Create(ctx context.Context, drone model.Drone) (model.Drone, error)
	Update(ctx context.Context, drone model.Drone) (model.Drone, error)
	Get(ctx context.Context, id uint64) (model.Drone, error)
	GetByAuthor(ctx context.Context, authorLogin string) ([]model.Drone, error)
	Delete(ctx context.Context, id uint64) error
}

type Publisher interface {
	Publish(ctx context.Context, event model.Event) (err error)
}

type DroneClient interface {
	StartDroneMission(ctx context.Context, drone model.Drone, mission model.Mission) (err error)
}
