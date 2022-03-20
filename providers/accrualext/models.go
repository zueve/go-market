package accrualext

var (
	StatusRegistered = "REGISTERED"
	StatusInvalid    = "INVALID"
	StatusProcessing = "PROCESSING"
	StatusProcessed  = "PROCESSED"
)

type Response struct {
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual"`
}
