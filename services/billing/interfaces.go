package billing

import "context"

type StorageExpected interface{
	NewWithdrawal(ctx context.Context, order WithdrawalOrder) (WithdrawalProcessedOrder, error)
	NewDeposit(ctx context.Context, order DepositOrder) (WithdrawalProcessedOrder, error)
	GetWithdrawalOrders(ctx context.Context, user string) ([WithdrawalProcessedOrder], error)
}
