package billing

import (
	"context"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services"

	"github.com/rs/zerolog"
)

type Service struct {
	Storage StorageExpected
}

func (s *Service) Process(ctx context.Context, order services.OrderValue) (services.ProcessedOrder, error) {
	s.log(ctx).Info().Msgf("Receive New order: %s - %d USD", order.Invoice, order.Amount)
	processedOrder, err := s.Storage.Process(ctx, order)
	if err != nil {
		return services.ProcessedOrder{}, err
	}
	return processedOrder, nil
}

func (s *Service) GetWithdrawalsOrders(ctx context.Context, user services.User) ([]services.ProcessedOrder, error) {
	orders, err := s.Storage.GetWithdrawalOrders(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Service) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "User").
		Str(logging.Layer, "service").
		Logger()
	return &logger
}
