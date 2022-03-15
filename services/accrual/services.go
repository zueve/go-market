package accrual

import (
	"context"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/billing"

	"github.com/rs/zerolog"
)

type Service struct {
	Storage StorageExpected
	Billing billing.Service
}

func (s *Service) NewOrder(ctx context.Context, user services.User, num int64) (OrderVal, error) {
	s.log(ctx).Info().Msgf("Receive New order: %d", num)
	order := OrderVal{Invoice: num, Status: StatusNew, UserID: user.ID}
	if err := s.Storage.NewOrder(ctx, order); err != nil {
		return order, err
	}
	return order, nil
}

func (s *Service) GetOrders(ctx context.Context, user services.User) ([]Order, error) {
	orders, err := s.Storage.GetOrders(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Service) UpdateOrderStatus(ctx context.Context, order OrderVal) (OrderVal, error) {
	if err := s.Storage.UpdateOrderStatus(ctx, order); err != nil {
		return OrderVal{}, err
	}
	return order, nil
}

func (s *Service) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "User").
		Str(logging.Layer, "service").
		Logger()
	return &logger
}
