package rest

type (
	AuthRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	AuthToken struct {
		AccessToken string `json:"token"`
	}
)
