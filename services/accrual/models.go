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
	Invoice int64  `json:"-"`
	Status  string `json:"status"`
	UserID  int    `json:"-"`
	Amount  int64  `json:"accrual"`
}

func (s *OrderVal) ToDeposit() services.OrderValue {
	return services.OrderValue{
		Invoice:   fmt.Sprintf("%d", s.Invoice),
		UserID:    s.UserID,
		Amount:    s.Amount * 100,
		IsDeposit: true,
	}
}

type Order struct {
	OrderVal
	InvoiceStr  string    `json:"number"`
	AmountFloat float32   `json:"accrual"`
	Created     time.Time `json:"uploaded_at"`
}
