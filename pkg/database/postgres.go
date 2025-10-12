package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPostgresDB(ctx context.Context, url string) (*PGDatabase, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &PGDatabase{Conn: conn, ctx: ctx}, nil
}

type PGDatabase struct {
	Conn *pgx.Conn
	ctx context.Context
}

func (db *PGDatabase) Close() error {
	return db.Conn.Close(db.ctx)
}