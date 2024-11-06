package repository

import "context"

type Command interface {
	Execute(ctx context.Context, query string, args ...interface{}) error
}
