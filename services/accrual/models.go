package accrual

const (
	StatusNew        string = "NEW"
	StatusProcessing string = "PROCESSING"
	StatusInvalid    string = "INVALID"
	StatusProcessed  string = "PROCESSED"
)

type AccrualOrder struct {
	Num    int64
	Status string
	User   string
}
