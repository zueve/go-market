package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"github.com/zueve/go-market/services/accrual"
)

func (s *Storage) GetUserOrders(ctx context.Context, userID int) ([]accrual.Order, error) {
	var operations []Accrual
	query := `
		SELECT *
		FROM accrual
		WHERE customer_id=$1
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
	var pgErr *pgconn.PgError
	// Add order
	query := `
		INSERT INTO accrual(customer_id, invoice, status, amount)
		VALUES(:userid, :invoice, :status, :amount)
	`
	if _, err := s.DB.NamedExecContext(ctx, query, &order); err != nil {
		if !(errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation) {
			return err
		}

		op, err := s.GetAccrualByInvoice(order.Invoice)
		if err != nil {
			return err
		}
		if op.CustomerID == order.UserID {
			return accrual.ErrOrderExist
		}
		return accrual.ErrInvoiceExist
	}
	return nil
}

func (s *Storage) UpdateOrderStatus(ctx context.Context, order accrual.OrderVal) error {
	query := `
		UPDATE accrual
		SET status=:status, amount=:amount, updated=now()
		WHERE invoice = :invoice
	`
	if _, err := s.DB.NamedExecContext(ctx, query, order); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAccrualByInvoice(invoice int64) (Accrual, error) {
	op := Accrual{}
	query := "SELECT * from accrual where invoice=$1"
	if err := s.DB.Get(&op, query, invoice); err != nil {
		return Accrual{}, err
	}
	return op, nil
}

func (s *Storage) GetOrders(ctx context.Context, status []string) ([]accrual.Order, error) {
	var operations []Accrual
	query := `
		SELECT *
		FROM accrual
		WHERE status=ANY($1)
		ORDER BY id desc
	`
	if err := s.DB.Select(&operations, query, status); err != nil {
		return nil, err
	}
	orders := make([]accrual.Order, len(operations))
	for i := range operations {
		orders[i] = operations[i].ToService()
	}
	return orders, nil
}
