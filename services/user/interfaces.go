package user

import "context"

type StorageExpected interface {
	CheckPassword(ctx context.Context, login string, password string) error
	Create(ctx context.Context, login string, password string) error
}
