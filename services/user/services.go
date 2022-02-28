package user

import (
	"context"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services"

	"github.com/rs/zerolog"
)

type UserService struct {
	Storage StorageExpected
}

func (s *UserService) Create(ctx context.Context, login string, password string) (services.User, error) {
	s.log(ctx).Info().Msgf("Create user: %s", login)
	if err := s.Storage.Create(ctx, login, password); err != nil {
		s.log(ctx).Info().Msgf("Fail %s", err)
		return services.User{}, err
	}
	return services.User{Login: login}, nil
}

func (s *UserService) Login(ctx context.Context, login string, password string) (services.User, error) {
	s.log(ctx).Info().Msgf("login as user: %s", login)
	if err := s.Storage.CheckPassword(ctx, login, password); err != nil {
		s.log(ctx).Info().Msgf("Fail %s", err)
		return services.User{}, err
	}
	return services.User{Login: login}, nil
}

func (s *UserService) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "User").
		Str(logging.Layer, "service").
		Logger()
	return &logger
}
