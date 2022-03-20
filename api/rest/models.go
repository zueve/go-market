package rest

import (
	"time"
)

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
	BalanceResponse struct {
		Balance   float32 `json:"balance"`
		Withdrawn float32 `json:"withdrawn"`
	}
	WithdrawalRequest struct {
		Invoice string  `json:"order"`
		Amount  float32 `json:"sum"`
	}

	WithdrawalOrder struct {
		Processed time.Time `json:"processed_at"`
		Invoice   string    `json:"order"`
		Amount    float32   `json:"sum"`
	}
)
