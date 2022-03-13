package rest

type (
	AuthRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	AuthToken struct {
		Token string `json:"access_token"`
		Type  string `json:"token_type"`
	}
)

type (
	WithdrawalRequest struct {
		Invoice string `json:"order"`
		Amount  int64  `json:"sum"`
	}
)
