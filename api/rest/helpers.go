package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services/user"

	"github.com/rs/zerolog"
)

func (s *Handler) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "rest").
		Str(logging.Layer, "api").
		Logger()
	return &logger
}

func (s *Handler) isValidRequest(
	ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{},
) bool {
	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeInternalError(ctx, w, err)
		return false
	}

	if err := json.Unmarshal(dataBytes, data); err != nil {
		httpError := NewValidationError(err.Error(), nil)
		s.writeHTTPError(ctx, w, httpError)
		return false
	}
	return true
}

func (s *Handler) writeResult(ctx context.Context, w http.ResponseWriter, status int, result interface{}) {
	response, err := json.Marshal(result)
	if err != nil {
		s.writeInternalError(ctx, w, err)
		return
	}
	s.writeResponse(ctx, w, status, response)
}

func (s *Handler) writeErr(ctx context.Context, w http.ResponseWriter, err error) {
	if httpError, ok := s.toHTTPError(err); ok {
		s.writeHTTPError(ctx, w, httpError)
	} else {
		s.writeInternalError(ctx, w, err)
	}
}

func (s *Handler) writeHTTPError(ctx context.Context, w http.ResponseWriter, httpError HTTPError) {
	data, err := httpError.ToJSON()
	if err != nil {
		s.writeInternalError(ctx, w, err)
	}

	s.writeResponse(ctx, w, httpError.GetStatusCode(), data)
}

func (s *Handler) writeInternalError(ctx context.Context, w http.ResponseWriter, err error) {
	s.log(ctx).Error().Err(err).Msg("InternalError")
	httpError := NewInternalError()
	s.writeHTTPError(ctx, w, httpError)
}

func (s *Handler) writeResponse(ctx context.Context, w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		s.log(ctx).Error().Err(err).Msg("could't write Response")
	}
}

func (s *Handler) toHTTPError(err error) (HTTPError, bool) {
	switch err {
	case user.ErrAuth:
		return NewAuthErr(err), true
	case user.ErrLoginExists:
		return NewLoginExistsErr(err), true
	default:
		return HTTPError{}, false
	}
}
