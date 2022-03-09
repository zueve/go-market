package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/billing"
)

func (s *Storage) Process(ctx context.Context, order services.OrderValue) (services.ProcessedOrder, error) {
	var pgErr *pgconn.PgError
	var billingID int

	// Get billing id
	query := "SELECT id from billing where customer_id=$1"
	if err := s.DB.GetContext(ctx, &billingID, query, order.UserID); err != nil {
		return services.ProcessedOrder{}, err
	}

	// Start transaction
	tx := s.DB.MustBegin()
	defer func() {
		if err := tx.Rollback(); err != nil {
			s.log(ctx).Error().Err(err).Msg("")
		}
	}()

	// Add order
	query = `
		INSERT INTO billing_order(balance_id, invoice, direction, amount)
		VALUES($1, $2, $3, $4)
	`
	if _, err := tx.ExecContext(ctx, query, billingID, order.Invoice, order.Direction(), order.Amount); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return services.ProcessedOrder{}, billing.ErrInvoiceExists
		}
		return services.ProcessedOrder{}, err
	}
	// Change balance
	var amount int64
	if order.IsDeposit {
		amount = order.Amount
	} else {
		amount = -order.Amount
	}
	query = `
		UPDATE billing SET amount = amount + $1
		WHERE id=$2
	`
	if _, err := tx.ExecContext(ctx, query, amount, billingID); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.CheckViolation {
			return services.ProcessedOrder{}, billing.ErrNotEnoughtMoney
		}
	}

	// Commit result
	if err := tx.Commit(); err != nil {
		return services.ProcessedOrder{}, err
	}

	return s.GetProcessedOrderByInvoice(order.Invoice)
}
func (s *Storage) GetWithdrawalOrders(ctx context.Context, user string) ([]services.ProcessedOrder, error) {
	return nil, nil
}

func (s *Storage) GetProcessedOrderByInvoice(invoice string) (services.ProcessedOrder, error) {
	order := services.ProcessedOrder{}
	query := "SELECT * from billing_order where invoice=$1"
	if err := s.DB.Get(&order, query, invoice); err != nil {
		return order, err
	}
	return order, nil
}
