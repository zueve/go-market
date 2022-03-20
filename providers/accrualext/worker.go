package accrualext

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/zueve/go-market/pkg/convert"
	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services/accrual"
)

type Worker struct {
	client  Client
	service Service
}

func (s *Worker) Loop(ctx context.Context, inpCh chan accrual.OrderVal) {
	for {
		select {
		case order := <-inpCh:
			s.Process(ctx, order, inpCh)
		case <-ctx.Done():
			return
		}
	}
}

func (s *Worker) Process(ctx context.Context, order accrual.OrderVal, inpCh chan accrual.OrderVal) {
	s.log(ctx).Info().Msgf("Start process order %d", order.Invoice)
	result, err := s.client.GetOrderStatus(ctx, order)
	if err != nil {
		s.log(ctx).Info().Msgf("On Error - send order back %d", order.Invoice)
		inpCh <- order
		return
	}
	switch result.Status {
	case StatusInvalid:
		order.Status = accrual.StatusInvalid
	case StatusProcessing:
		order.Status = accrual.StatusProcessing
	case StatusProcessed:
		order.Status = accrual.StatusProcessed
		order.Amount = convert.MoneyToMinor(result.Accrual)
	default:
		s.log(ctx).Info().Msgf("On Error - send order back %d", order.Invoice)
		inpCh <- order
		return
	}

	s.log(ctx).Info().Msgf("Update order status %d", order.Invoice)
	if _, err = s.service.UpdateOrderStatus(ctx, order); err != nil {
		s.log(ctx).Err(err).Msg("")
		return
	}
}

func (s *Worker) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "accrual.Worker").
		Str(logging.Layer, "providers").
		Logger()
	return &logger
}
