package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	// "github.com/golang-migrate/migrate"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zueve/go-market/api/rest"
	"github.com/zueve/go-market/config"
	"github.com/zueve/go-market/providers/postgres"
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

	if err = initLogger(conf); err != nil {
		return err
	}

	storage := postgres.Storage{}
	userSrv := user.UserService{Storage: &storage}
	tokenAuth := jwtauth.New("HS256", []byte(conf.Secret), nil)
	handler, err := rest.New(userSrv, tokenAuth)
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

func initLogger(config *config.Config) error {
	level, err := zerolog.ParseLevel(strings.ToLower(config.LogLevel))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    !config.LogColor,
		},
	)
	return nil
}
