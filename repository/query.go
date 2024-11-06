package repository

import "context"

type Query interface {
	Query(ctx context.Context, query string, args ...interface{}) (interface{}, error)
}
