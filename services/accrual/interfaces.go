package accrual

import "context"

type StorageExpected interface {
	NewOrder(ctx context.Context, order OrderVal) error
	GetOrders(ctx context.Context, user string) ([]Order, error)
	UpdateOrderStatus(ctx context.Context, order OrderVal) error
}
