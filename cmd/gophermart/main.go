package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/jwtauth/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/zueve/go-market/api/rest"
	"github.com/zueve/go-market/config"
	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/pkg/migrate"
	"github.com/zueve/go-market/providers/postgres"
	"github.com/zueve/go-market/services/billing"
	"github.com/zueve/go-market/services/user"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("Execution failed")
		os.Exit(1)
	}
}

func run() error {
	conf, err := config.NewFromEnv()
	if err != nil {
		return err
	}

	if err = logging.InitLogger(conf.LogLevel, conf.LogColor); err != nil {
		return err
	}

	// run migrations on startup
	if err = migrate.Run(conf.DatabaseDSN, conf.MigrateFile); err != nil {
		return err
	}

	db, err := sqlx.Open("pgx", conf.DatabaseDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	storage := postgres.Storage{DB: db}
	userSrv := user.Service{Storage: &storage}
	billingSrv := billing.Service{Storage: &storage}
	tokenAuth := jwtauth.New("HS256", []byte(conf.Secret), nil)
	handler, err := rest.New(userSrv, billingSrv, tokenAuth)
	if err != nil {
		return err
	}

	httpServer := http.Server{
		Addr:    conf.ServerAddress,
		Handler: handler.Handler,
	}

	stop := make(chan os.Signal, 1)
	serverError := make(chan error, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		serverError <- httpServer.ListenAndServe()
	}()
	log.Info().Msgf("Starting at %s...", conf.ServerAddress)

	select {
	case <-stop:
	case err = <-serverError:
	}

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		return errors.New("unexpected err on graceful shutdown")
	}
	return nil
}
