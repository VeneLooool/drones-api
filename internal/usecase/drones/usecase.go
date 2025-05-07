package drones

import (
	"context"
	"errors"

	"github.com/VeneLooool/drones-api/internal/model"
	"github.com/VeneLooool/drones-api/internal/pkg/error_hub"
	"github.com/jackc/pgx/v4"
)

type UseCase struct {
	repo Repo

	publisher   Publisher
	droneClient DroneClient
}

func New(repo Repo, publisher Publisher, droneClient DroneClient) *UseCase {
	return &UseCase{
		repo:        repo,
		publisher:   publisher,
		droneClient: droneClient,
	}
}

func (u *UseCase) Create(ctx context.Context, drone model.Drone) (model.Drone, error) {
	drone.SetDefaultStatus()

	return u.repo.Create(ctx, drone)
}

func (u *UseCase) Update(ctx context.Context, new model.Drone) (model.Drone, error) {
	old, err := u.Get(ctx, new.ID)
	if err != nil {
		return model.Drone{}, err
	}

	new, err = u.repo.Update(ctx, new)
	if err != nil {
		return model.Drone{}, err
	}

	if old.Status != new.Status {
		err = u.publisher.Publish(ctx, model.Event{
			EventType: model.EventTypeDroneChangeStatus,
			Drone:     new,
		})
		if err != nil {
			return model.Drone{}, err
		}
	}
	return new, nil
}

func (u *UseCase) Get(ctx context.Context, id uint64) (model.Drone, error) {
	drone, err := u.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Drone{}, error_hub.ErrDroneNotFound
		}
		return model.Drone{}, err
	}
	return drone, nil
}

func (u *UseCase) GetByAuthor(ctx context.Context, authorLogin string) ([]model.Drone, error) {
	return u.repo.GetByAuthor(ctx, authorLogin)
}

func (u *UseCase) Delete(ctx context.Context, id uint64) error {
	return u.repo.Delete(ctx, id)
}

func (u *UseCase) StartDroneMission(ctx context.Context, droneID uint64, mission model.Mission) error {
	drone, err := u.Get(ctx, droneID)
	if err != nil {
		return err
	}

	if drone.Status != model.DroneStatusAvailable {
		return error_hub.ErrDroneNotAvailable
	}

	if err = u.droneClient.StartDroneMission(ctx, drone, mission); err != nil {
		return err
	}
	return nil
}
