package billing

import "context"

type StorageExpected interface {
	NewWithdrawal(ctx context.Context, order WithdrawalOrder) (WithdrawalProcessedOrder, error)
	NewDeposit(ctx context.Context, order DepositOrder) (DepositProcessedOrder, error)
	GetWithdrawalOrders(ctx context.Context, user string) ([]WithdrawalProcessedOrder, error)
}
