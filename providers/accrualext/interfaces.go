package accrualext

import (
	"context"

	"github.com/zueve/go-market/services/accrual"
)

type Client interface {
	GetOrderStatus(ctx context.Context, order accrual.OrderVal) (Response, error)
}

type Service interface {
	UpdateOrderStatus(ctx context.Context, order accrual.OrderVal) (accrual.OrderVal, error)
}
