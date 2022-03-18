package accrualext

var (
	StatusRegistered = "REGISTERED"
	StatusInvalid    = "INVALID"
	StatusProcessing = "PROCESSING"
	StatusProcessed  = "PROCESSED"
)

type Response struct {
	Status  string `json:"status"`
	Accrual int64  `json:"accrual"`
}
