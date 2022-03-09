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

func (s *Service) NewWithdrawal(ctx context.Context, order WithdrawalOrder) (WithdrawalProcessedOrder, error) {
	s.log(ctx).Info().Msgf("Receive New order: %d - %d USD", order.Invoice, order.Amount)
	processedOrder, err := s.Storage.NewWithdrawal(ctx, order)
	if err != nil {
		return WithdrawalProcessedOrder{}, err
	}
	return processedOrder, nil
}

func (s *Service) GetWithdrawalsOrders(ctx context.Context, user services.User) ([]WithdrawalProcessedOrder, error) {
	orders, err := s.Storage.GetWithdrawalOrders(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Service) NewDeposit(ctx context.Context, order DepositOrder) (DepositProcessedOrder, error) {
	processedOrder, err := s.Storage.NewDeposit(ctx, order)

	if err != nil {
		return DepositProcessedOrder{}, err
	}
	return processedOrder, nil
}

func (s *Service) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "User").
		Str(logging.Layer, "service").
		Logger()
	return &logger
}
