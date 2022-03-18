package accrualext

import (
	"context"

	"github.com/zueve/go-market/services/accrual"
)

type Provider struct {
	inpCh chan accrual.OrderVal
}

func New(
	ctx context.Context, client Client, service Service, workerNum int, inpCh chan accrual.OrderVal,
) (*Provider, error) {
	for i := 0; i < workerNum; i++ {
		w := Worker{client: client, service: service}
		go w.Loop(ctx, inpCh)
	}
	return &Provider{inpCh: inpCh}, nil
}

func (s *Provider) ProcessOrder(ctx context.Context, order accrual.OrderVal) error {
	s.inpCh <- order
	return nil
}
