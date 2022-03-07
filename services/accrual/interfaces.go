package accrual

import "context"

type StorageExpected interface {
	NewOrder(ctx context.Context, order AccrualOrder) error
	GetOrders(ctx context.Context, user string) ([]AccrualOrder, error)
	UpdateOrderStatus(ctx context.Context, order AccrualOrder) error
}
