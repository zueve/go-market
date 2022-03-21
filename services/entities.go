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
		Invoice   string `json:"order"`
		UserID    int    `json:"-"`
		Amount    int64  `json:"sum"`
		IsDeposit bool   `json:"is_deposit"`
	}
	ProcessedOrder struct {
		OrderValue
		ID        int       `json:"-"`
		Processed time.Time `json:"processed_at"`
	}
)

func (s *OrderValue) Direction() string {
	if s.IsDeposit {
		return DirectionDeposit
	}
	return DirectionWithdrawal
}
