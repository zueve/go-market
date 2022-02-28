package postgres

import "context"

type Storage struct{}

func (s *Storage) Create(ctx context.Context, login string, password string) error {
	return nil
}
func (s *Storage) CheckPassword(ctx context.Context, login string, password string) error {
	return nil
}
