package services

import "time"

type User struct {
	Login string
	ID    int
}

type (
	OrderValue struct {
		Invoice   string
		User      string
		Amount    int64
		IsDeposit bool
	}
	ProcessedOrder struct {
		OrderValue
		ID        int
		Processed time.Time
	}
)
