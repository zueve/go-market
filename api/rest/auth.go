package rest

import (
	"context"
	"net/http"

	"github.com/zueve/go-market/services"
)

func (s *Handler) register(w http.ResponseWriter, r *http.Request) {
	var data AuthRequest
	ctx := r.Context()
	if !s.isValidRequest(ctx, w, r, &data) {
		return
	}
	u, err := s.UserService.Create(ctx, data.Login, data.Password)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	s.writeToken(ctx, w, u)
}

func (s *Handler) login(w http.ResponseWriter, r *http.Request) {
	var data AuthRequest
	ctx := r.Context()
	if !s.isValidRequest(ctx, w, r, &data) {
		return
	}
	u, err := s.UserService.Login(ctx, data.Login, data.Password)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	s.writeToken(ctx, w, u)

}

func (s *Handler) writeToken(ctx context.Context, w http.ResponseWriter, u services.User) {
	// create auth token
	_, token, err := s.TokenAuth.Encode(map[string]interface{}{"user": u})
	if err != nil {
		s.writeInternalError(ctx, w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/plain")
	if _, err := w.Write([]byte(token)); err != nil {
		s.log(ctx).Error().Err(err).Msg("")
	}
}
