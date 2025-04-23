package db

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DataBaseAdapter struct {
	conn *pgx.Conn
}

func New(ctx context.Context) (*DataBaseAdapter, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &DataBaseAdapter{
		conn: conn,
	}, nil
}

func (d *DataBaseAdapter) Conn() *pgx.Conn {
	return d.conn
}

func (d *DataBaseAdapter) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return d.conn.Exec(ctx, sql, arguments...)
}

func (d *DataBaseAdapter) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	return d.conn.Query(ctx, sql, arguments...)
}

func (d *DataBaseAdapter) Close(ctx context.Context) error {
	return d.conn.Close(ctx)
}
