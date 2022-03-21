package rest

import (
	"encoding/json"
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
		Balance   json.Number `json:"current"`
		Withdrawn json.Number `json:"withdrawn"`
	}
	WithdrawalRequest struct {
		Invoice string  `json:"order"`
		Amount  float32 `json:"sum"`
	}

	WithdrawalOrder struct {
		Processed time.Time   `json:"processed_at"`
		Invoice   string      `json:"order"`
		Amount    json.Number `json:"sum"`
	}
)

type (
	AccrualOrder struct {
		Invoice string      `json:"number"`
		Amount  json.Number `json:"accrual"`
		Created time.Time   `json:"uploaded_at"`
		Status  string      `json:"status"`
	}
)
