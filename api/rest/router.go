package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services/billing"
	"github.com/zueve/go-market/services/user"
)

type Handler struct {
	Handler        chi.Router
	UserService    user.Service
	BillingService billing.Service
	TokenAuth      *jwtauth.JWTAuth
}

func New(userSrv user.Service, billingSrv billing.Service, tokenAuth *jwtauth.JWTAuth) (Handler, error) {
	router := chi.NewRouter()
	h := Handler{
		Handler:        router,
		UserService:    userSrv,
		BillingService: billingSrv,
		TokenAuth:      tokenAuth,
	}
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Recoverer)
	router.Use(logging.AccessLog())

	router.Post("/api/user/register", h.register)
	router.Post("/api/user/login", h.login)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/api/user/balance", h.getBalance)
		r.Get("/api/user/balance/withdraw", h.getWithdrawals)
		r.Post("/api/user/balance/withdraw", h.createWithdrawal)
	})

	return h, nil
}
