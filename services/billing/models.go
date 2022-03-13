package billing

type Balance struct {
	Balance   int64 `json:"balance"`
	Withdrawn int64 `json:"withdrawn"`
}
