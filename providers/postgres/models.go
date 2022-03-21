package postgres

import (
	"time"

	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/accrual"
)

type Operation struct {
	ID         int       `db:"id"`
	Invoice    string    `db:"invoice"`
	BillingID  int       `db:"billing_id"`
	CustomerID int       `db:"customer_id"`
	Direction  string    `db:"direction"`
	Amount     int64     `db:"amount"`
	Created    time.Time `db:"created"`
}

func (s *Operation) IsDeposit() bool {
	return s.Direction == services.DirectionDeposit
}

func (s *Operation) ToOrder() services.ProcessedOrder {
	return services.ProcessedOrder{
		OrderValue: services.OrderValue{
			Invoice:   s.Invoice,
			Amount:    s.Amount,
			UserID:    s.CustomerID,
			IsDeposit: s.IsDeposit(),
		},
		ID:        s.ID,
		Processed: s.Created,
	}
}

type Accrual struct {
	ID         int       `db:"id"`
	CustomerID int       `db:"customer_id"`
	Invoice    int64     `db:"invoice"`
	Amount     int64     `db:"amount"`
	Status     string    `db:"status"`
	Created    time.Time `db:"created"`
	Updated    time.Time `db:"updated"`
}

func (s *Accrual) ToService() accrual.Order {
	return accrual.Order{
		OrderVal: accrual.OrderVal{
			Invoice: s.Invoice,
			UserID:  s.CustomerID,
			Status:  s.Status,
			Amount:  s.Amount,
		},
		Created: s.Created,
	}
}
