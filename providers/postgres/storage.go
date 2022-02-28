package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/zueve/go-market/services/user"
)

type Storage struct {
	DB *sqlx.DB
}

func (s *Storage) Create(ctx context.Context, login string, password string) error {
	query := "INSERT INTO customer(login, password_hash) VALUES($1, $2)"

	var pgErr *pgconn.PgError
	if _, err := s.DB.ExecContext(ctx, query, login, password); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return user.LoginExistsErr
		} else if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) CheckPassword(ctx context.Context, login string, password string) error {
	query := "SELECT password_hash from customer where login = $1"
	var stored_hash string
	if err := s.DB.GetContext(ctx, &stored_hash, query, login); err != nil {
		if err == sql.ErrNoRows {
			return user.AuthErr
		}
		return err
	}
	if stored_hash != password {
		return user.AuthErr
	}
	return nil
}
