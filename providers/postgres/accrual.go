package postgres

import (
	"context"

	"github.com/zueve/go-market/services/accrual"
)

func (s *Storage) GetOrders(ctx context.Context, userID int) ([]accrual.Order, error) {
	var operations []Accrual
	query := `
		SELECT *
		JOIN accrual
		WHERE b.customer_id=$1
		ORDER BY id desc
	`
	if err := s.DB.Select(&operations, query, userID); err != nil {
		return nil, err
	}
	orders := make([]accrual.Order, len(operations))
	for i := range operations {
		orders[i] = operations[i].ToService()
	}
	return orders, nil
}

func (s *Storage) NewOrder(ctx context.Context, order accrual.OrderVal) error {
	// Add order
	query := `
		INSERT INTO accrual(customer_id, invoice, status, amount)
		VALUES(:UserID, :Invoice, :Status, :Amount)
		ON CONFLICT DO NOTHING
	`
	if _, err := s.DB.NamedExecContext(ctx, query, order); err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateOrderStatus(ctx context.Context, order accrual.OrderVal) error {
	query := `
		UPDATE accrual
		SET status=:Status, amount:Amount, updated=now()
		WHERE invoice = :Invoice
	`
	if _, err := s.DB.NamedExecContext(ctx, query, order); err != nil {
		return err
	}
	return nil
}
