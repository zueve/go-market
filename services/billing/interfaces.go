package billing

import (
	"context"

	"github.com/zueve/go-market/services"
)

type StorageExpected interface {
	Process(ctx context.Context, order services.OrderValue) (services.ProcessedOrder, error)
	GetWithdrawalOrders(ctx context.Context, user string) ([]services.ProcessedOrder, error)
}
