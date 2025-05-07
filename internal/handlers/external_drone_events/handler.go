package external_drone_events

import (
	"context"
	"errors"

	"github.com/VeneLooool/drones-api/internal/model"
	"github.com/VeneLooool/drones-api/internal/pkg/error_hub"
)

type Handler struct {
	droneUC droneUC
}

func New(droneUC droneUC) *Handler {
	return &Handler{
		droneUC: droneUC,
	}
}

func (h *Handler) Handle(ctx context.Context, event model.Event) error {
	if event.EventType != model.EventTypeDroneChangeStatus {
		return nil
	}

	drone, err := h.droneUC.Get(ctx, event.Drone.ID)
	if err != nil && !errors.Is(err, error_hub.ErrDroneNotFound) {
		return err
	}
	drone.SetStatus(event.Drone.Status)

	_, err = h.droneUC.Update(ctx, drone)
	if err != nil && !errors.Is(err, error_hub.ErrDroneNotFound) {
		return err
	}
	return nil
}
