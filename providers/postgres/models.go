package postgres

import (
	"time"

	"github.com/zueve/go-market/services"
)

type Operation struct {
	ID        int       `db:"id"`
	Invoice   string    `db:"invoice"`
	BalanceID int       `db:"balance_id"`
	UserID    int       `db:"user_id"`
	Direction string    `db:"direction"`
	Amount    int64     `db:"amount"`
	Created   time.Time `db:"created"`
}

func (s *Operation) IsDeposit() bool {
	return s.Direction == "DEPOSIT"
}

func (s *Operation) ToOrder() services.ProcessedOrder {
	return services.ProcessedOrder{
		OrderValue: services.OrderValue{
			Invoice:   s.Invoice,
			Amount:    s.Amount,
			UserID:    s.UserID,
			IsDeposit: s.IsDeposit(),
		},
		ID:        s.ID,
		Processed: s.Created,
	}
}
