package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"

	"github.com/zueve/go-market/services/user"
)

type Handler struct {
	Handler     chi.Router
	UserService user.UserService
	TokenAuth   *jwtauth.JWTAuth
}

func New(service user.UserService, tokenAuth *jwtauth.JWTAuth) (Handler, error) {
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

	router.Post("/api/register", h.register)
	router.Post("/api/login", h.login)
	return h, nil
}
