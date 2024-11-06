package service

import (
	"context"
	"lambda-golang-addons-boilerplate/repository"
)

type Service struct {
	command repository.Command
	query   repository.Query
	cache   repository.Cache
}

func NewService(command repository.Command, query repository.Query, cache repository.Cache) *Service {
	return &Service{command: command, query: query, cache: cache}
}

// Command
func (s *Service) CreateItem(ctx context.Context, item string) error {
	return s.command.Execute(ctx, "INSERT INTO items (name) VALUES ($1)", item)
}

// Query
func (s *Service) GetItem(ctx context.Context, id string) (interface{}, error) {
	return s.query.Query(ctx, "SELECT * FROM items WHERE id=$1", id)
}
