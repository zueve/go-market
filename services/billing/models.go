package billing

import "time"

type (
	WithdrawalOrder struct {
		Invoice string
		User    string
		Amount  int64
	}

	WithdrawalProcessedOrder struct {
		WithdrawalOrder
		processed time.Time
		ID        int64
	}
)

type (
	DepositOrder struct {
		Invoice string
		User    string
		Amount  int64
	}

	DepositProcessedOrder struct {
		DepositOrder
		processed time.Time
	}
)
