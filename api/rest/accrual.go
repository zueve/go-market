package rest

import (
	"fmt"
	"net/http"
)

func (s *Handler) getAccrualOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := s.getUser(ctx)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}

	orders, err := s.AccrualService.GetOrders(ctx, user)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}
	if len(orders) == 0 {
		s.writeResult(ctx, w, http.StatusNoContent, orders)
		return
	}
	s.writeResult(ctx, w, http.StatusOK, orders)
}

func (s *Handler) createAccrualOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := s.getUser(ctx)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}

	var num int64
	if !s.isValidRequest(ctx, w, r, &num) {
		return
	}
	if !s.isValidInvoice(ctx, w, fmt.Sprint(num)) {
		return
	}

	orders, err := s.AccrualService.NewOrder(ctx, user, num)
	if err != nil {
		s.writeErr(ctx, w, err)
		return
	}

	s.writeResult(ctx, w, http.StatusAccepted, orders)
}
