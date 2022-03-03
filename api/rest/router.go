package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services/user"
)

type Handler struct {
	Handler     chi.Router
	UserService user.Service
	TokenAuth   *jwtauth.JWTAuth
}

func New(service user.Service, tokenAuth *jwtauth.JWTAuth) (Handler, error) {
	router := chi.NewRouter()
	h := Handler{
		Handler:     router,
		UserService: service,
		TokenAuth:   tokenAuth,
	}
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Recoverer)
	router.Use(logging.AccessLog())

	router.Post("/api/register", h.register)
	router.Post("/api/login", h.login)
	return h, nil
}
