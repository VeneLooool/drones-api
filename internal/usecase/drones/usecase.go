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
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, drone model.Drone) (model.Drone, error) {
	drone.SetDefaultStatus()

	return u.repo.Create(ctx, drone)
}

func (u *UseCase) Update(ctx context.Context, drone model.Drone) (model.Drone, error) {
	return u.repo.Update(ctx, drone)
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

func (u *UseCase) Delete(ctx context.Context, id uint64) error {
	return u.repo.Delete(ctx, id)
}
