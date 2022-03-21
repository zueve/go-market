package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/user"
)

var _ user.StorageExpected = (*Storage)(nil)

type Storage struct {
	DB *sqlx.DB
}

func (s *Storage) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "User").
		Str(logging.Layer, "service").
		Logger()
	return &logger
}

func (s *Storage) Create(ctx context.Context, login string, password string) (services.User, error) {
	var id int
	var pgErr *pgconn.PgError

	tx := s.DB.MustBegin()
	defer func() {
		if err := tx.Rollback(); err != nil {
			s.log(ctx).Error().Err(err).Msg("")
		}
	}()
	// Create user
	query := "INSERT INTO customer(login, password_hash) VALUES($1, $2) returning id"

	if err := tx.GetContext(ctx, &id, query, login, password); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return services.User{}, user.ErrLoginExists
		}
		return services.User{}, err
	}

	// Create billind
	query = "INSERT INTO billing(customer_id, amount) VALUES($1, $2)"
	if _, err := tx.ExecContext(ctx, query, id, 0); err != nil {
		return services.User{}, err
	}
	// Commit result
	if err := tx.Commit(); err != nil {
		return services.User{}, err
	}

	return services.User{Login: login, ID: id}, nil
}

func (s *Storage) CheckPassword(ctx context.Context, login string, password string) (services.User, error) {
	query := "SELECT password_hash, id from customer where login = $1"

	type Row struct {
		ID         int    `db:"id"`
		StoredHash string `db:"password_hash"`
	}
	var row Row
	if err := s.DB.GetContext(ctx, &row, query, login); err != nil {
		if err == sql.ErrNoRows {
			return services.User{}, user.ErrAuth
		}
		return services.User{}, err
	}
	if row.StoredHash != password {
		return services.User{}, user.ErrAuth
	}
	return services.User{Login: login, ID: row.ID}, nil
}
