package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/user"
)

var _ user.StorageExpected = (*Storage)(nil)

type Storage struct {
	DB *sqlx.DB
}

func (s *Storage) Create(ctx context.Context, login string, password string) (services.User, error) {
	query := "INSERT INTO customer(login, password_hash) VALUES($1, $2) returning id"

	var (
		id    int
		pgErr *pgconn.PgError
	)
	if err := s.DB.GetContext(ctx, &id, query, login, password); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return services.User{}, user.ErrLoginExists
		}
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
