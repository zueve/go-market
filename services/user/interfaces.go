package user

import (
	"context"

	"github.com/zueve/go-market/services"
)

type StorageExpected interface {
	CheckPassword(ctx context.Context, login string, password string) (services.User, error)
	Create(ctx context.Context, login string, password string) (services.User, error)
}
