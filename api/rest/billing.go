package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/zueve/go-market/services"
)

func (s *Handler) getWithdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := s.getUser(ctx)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}

	orders, err := s.BillingService.GetWithdrawalsOrders(ctx, user)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	s.writeResult(ctx, w, http.StatusOK, orders)
}

func (s *Handler) getUser(ctx context.Context) (services.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return services.User{}, err
	}

	userBites, err := json.Marshal(claims["user"])
	if err != nil {
		return services.User{}, err
	}

	var user services.User
	if err = json.Unmarshal(userBites, &user); err != nil {
		return services.User{}, err
	}
	return user, nil
}
