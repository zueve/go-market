package logging

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	RequestIDKey = "req_id"
	Source       = "Source"
	Layer        = "Layer"
)

func InitLogger(level string, withColor bool) error {
	lvl, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(lvl)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    !withColor,
		},
	)
	return nil
}

func GetLogger(ctx context.Context) zerolog.Logger {
	requestID := middleware.GetReqID(ctx)
	return log.With().Str(RequestIDKey, requestID).Logger()
}
