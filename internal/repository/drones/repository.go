package drones

import (
	"context"

	"github.com/VeneLooool/drones-api/internal/model"
	"github.com/VeneLooool/drones-api/internal/pkg/db"
	"github.com/VeneLooool/drones-api/internal/pkg/ql"
	common "github.com/VeneLooool/drones-api/internal/repository"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/huandu/go-sqlbuilder"
)

type Repo struct {
	db db.DataBase
}

func New(db db.DataBase) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, drone model.Drone) (model.Drone, error) {
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto(Table).Cols(
		Name.Short(),
		Status.Short(),
	).Values(
		drone.Name,
		drone.Status,
	)
	query, args := common.ReturningAll(ib).Build()

	var result model.Drone
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Drone{}, err
	}
	return result, nil
}

func (r *Repo) Update(ctx context.Context, drone model.Drone) (model.Drone, error) {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(Table)
	ub = ub.Set(ql.Fields{Name, Status}.ToAssignments(ub, drone.Name, drone.Status)...)
	ub = ub.Where(ub.Equal(ID.Short(), drone.ID))
	query, args := common.ReturningAll(ub).Build()

	var result model.Drone
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Drone{}, err
	}
	return result, nil
}

func (r *Repo) Get(ctx context.Context, id uint64) (model.Drone, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(ID.Short(), id))
	query, args := sb.Build()

	var result model.Drone
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Drone{}, err
	}
	return result, nil
}

func (r *Repo) Delete(ctx context.Context, id uint64) error {
	db := sqlbuilder.PostgreSQL.NewDeleteBuilder().DeleteFrom(Table)
	query, args := db.Where(db.Equal(ID.Short(), id)).Build()

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
