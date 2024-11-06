package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresQuery struct {
	pool *pgxpool.Pool
}

func NewPostgresQuery(connString string) (*PostgresQuery, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresQuery{pool: pool}, nil
}

func (p *PostgresQuery) Query(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Process rows and return result
	return nil, nil
}
