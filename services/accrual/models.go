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
}

type Order struct {
	OrderVal
	Amount  int64     `json:"accrual"`
	Created time.Time `json:"uploaded_at"`
}
