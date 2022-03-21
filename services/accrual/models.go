package accrual

import (
	"fmt"
	"time"

	"github.com/zueve/go-market/services"
)

const (
	StatusNew        string = "NEW"
	StatusProcessing string = "PROCESSING"
	StatusInvalid    string = "INVALID"
	StatusProcessed  string = "PROCESSED"
)

type OrderVal struct {
	Invoice int64
	Status  string
	UserID  int
	Amount  int64
}

func (s *OrderVal) ToDeposit() services.OrderValue {
	return services.OrderValue{
		Invoice:   fmt.Sprintf("%d", s.Invoice),
		UserID:    s.UserID,
		Amount:    s.Amount,
		IsDeposit: true,
	}
}

type Order struct {
	OrderVal
	Created time.Time
}
