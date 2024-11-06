package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresCommand struct {
	pool *pgxpool.Pool
}

func NewPostgresCommand(connString string) (*PostgresCommand, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresCommand{pool: pool}, nil
}

func (p *PostgresCommand) Execute(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.pool.Exec(ctx, query, args...)
	return err
}
