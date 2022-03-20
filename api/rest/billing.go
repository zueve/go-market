package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/user"
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
	if len(orders) == 0 {
		s.writeResult(ctx, w, http.StatusNoContent, orders)
		return
	}
	s.writeResult(ctx, w, http.StatusOK, ToWithdrawalResponse(orders))
}

func (s *Handler) getBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := s.getUser(ctx)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}

	balance, err := s.BillingService.GetBalance(ctx, user)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	s.writeResult(ctx, w, http.StatusOK, ToBalanceResponse(&balance))
}

func (s *Handler) createWithdrawal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := s.getUser(ctx)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	var request WithdrawalRequest
	if !s.isValidRequest(ctx, w, r, &request) {
		return
	}
	order := services.OrderValue{
		Invoice:   request.Invoice,
		UserID:    user.ID,
		Amount:    MoneyToMinor(request.Amount),
		IsDeposit: false,
	}

	balance, err := s.BillingService.Process(ctx, order)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	s.writeResult(ctx, w, http.StatusOK, balance)
}

func (s *Handler) getUser(ctx context.Context) (services.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return services.User{}, err
	}
	if len(claims) == 0 {
		return services.User{}, user.ErrAuth
	}

	userBites, err := json.Marshal(claims["user"])
	if err != nil {
		return services.User{}, err
	}

	var user services.User
	if err = json.Unmarshal(userBites, &user); err != nil {
		return services.User{}, err
	}
	s.log(ctx).Info().Msgf("login as user %s, id=%d", user.Login, user.ID)
	return user, nil
}
