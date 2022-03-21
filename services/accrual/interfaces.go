package accrual

import (
	"context"

	"github.com/zueve/go-market/services"
)

type StorageExpected interface {
	NewOrder(ctx context.Context, order OrderVal) error
	GetUserOrders(ctx context.Context, userID int) ([]Order, error)
	GetOrders(ctx context.Context, status []string) ([]Order, error)
	UpdateOrderStatus(ctx context.Context, order OrderVal) error
}

type ExternalAccrualExpected interface {
	ProcessOrder(ctx context.Context, order OrderVal) error
}

type BillingService interface {
	Process(ctx context.Context, order services.OrderValue) (services.ProcessedOrder, error)
}
