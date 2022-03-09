package accrual

import "context"

type StorageExpected interface {
	NewOrder(ctx context.Context, order Order) error
	GetOrders(ctx context.Context, user string) ([]Order, error)
	UpdateOrderStatus(ctx context.Context, order Order) error
}
