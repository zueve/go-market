package services

import "time"

const (
	DirectionDeposit    = "DEPOSIT"
	DirectionWithdrawal = "WITHDRAWAL"
)

type User struct {
	Login string
	ID    int
}

type (
	OrderValue struct {
		Invoice   string
		UserID    int
		Amount    int64
		IsDeposit bool
	}
	ProcessedOrder struct {
		OrderValue
		ID        int
		Processed time.Time
	}
)

func (s *OrderValue) Direction() string {
	if s.IsDeposit {
		return DirectionDeposit
	}
	return DirectionWithdrawal
}
