package logging

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func AccessLog() func(http.Handler) http.Handler {
	consumer := hlog.NewHandler(log.Logger)
	producer := hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		logger := GetLogger(r.Context())
		logger.Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	})
	return func(next http.Handler) http.Handler {
		return consumer(producer(next))
	}
}
