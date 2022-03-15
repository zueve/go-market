package accrual

import "time"

const (
	StatusNew        string = "NEW"
	StatusProcessing string = "PROCESSING"
	StatusInvalid    string = "INVALID"
	StatusProcessed  string = "PROCESSED"
)

type OrderVal struct {
	Invoice int64  `json:"number"`
	Status  string `json:"status"`
	UserID  int    `json:"-"`
	Amount  int64  `json:"accrual"`
}

type Order struct {
	OrderVal
	Created time.Time `json:"uploaded_at"`
}
