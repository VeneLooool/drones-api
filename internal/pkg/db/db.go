package db

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DataBaseAdapter struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context) (*DataBaseAdapter, error) {
	pool, err := pgxpool.Connect(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	return &DataBaseAdapter{
		pool: pool,
	}, nil
}

func (d *DataBaseAdapter) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, sql, arguments...)
}

func (d *DataBaseAdapter) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	return d.pool.Query(ctx, sql, arguments...)
}

func (d *DataBaseAdapter) Close(ctx context.Context) error {
	d.pool.Close()
	return nil
}
