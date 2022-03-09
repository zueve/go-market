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
		ID        int
		Processed time.Time
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
		ID        int
		Processed time.Time
	}
)
