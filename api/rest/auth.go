package rest

import (
	"net/http"

	"github.com/zueve/go-market/services/user"
)

func (s *Handler) register(w http.ResponseWriter, r *http.Request) {
	var data AuthRequest
	ctx := r.Context()
	if !s.isValidRequest(ctx, w, r, &data) {
		return
	}
	u, err := s.UserService.Create(ctx, data.Login, data.Password)
	if err != nil {
		if err == user.LoginExistsErr {
			httpErr := NewLoginExistsErr(err)
			s.writeHTTPError(ctx, w, httpErr)
			return
		}
		s.writeInternalError(ctx, w, err)
	}
	s.writeResult(ctx, w, http.StatusOK, u)
}

func (s *Handler) login(w http.ResponseWriter, r *http.Request) {
	var data AuthRequest
	ctx := r.Context()
	if !s.isValidRequest(ctx, w, r, &data) {
		return
	}
	u, err := s.UserService.Login(ctx, data.Login, data.Password)
	if err != nil {
		if err == user.AuthErr {
			httpErr := NewAuthErr(err)
			s.writeHTTPError(ctx, w, httpErr)
			return
		}
		s.writeInternalError(ctx, w, err)
	}
	// create token
	_, token, err := s.TokenAuth.Encode(map[string]interface{}{"user": u})
	if err != nil {
		s.writeInternalError(ctx, w, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/plain")
	if _, err := w.Write([]byte(token)); err != nil {
		s.log(ctx).Error().Err(err).Msg("")
	}
}
