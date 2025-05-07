package external_drone_events

import (
	"context"

	"github.com/VeneLooool/drones-api/internal/model"
)

type droneUC interface {
	Update(ctx context.Context, new model.Drone) (model.Drone, error)
	Get(ctx context.Context, id uint64) (model.Drone, error)
}
