package rest

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"net/http"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services"
	"github.com/zueve/go-market/services/billing"
	"github.com/zueve/go-market/services/user"

	"github.com/ShiraazMoollatjie/goluhn"
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
	case billing.ErrNotEnoughtMoney:
		return NewOutOfMoney(err), true
	default:
		return HTTPError{}, false
	}
}

func MoneyToMinor(v float32) int64 {
	return int64(math.Round(float64(v * 100.0)))
}

func MinorToMoney(v int64) float32 {
	return float32(v) / 100
}

func ToBalanceResponse(s *billing.Balance) BalanceResponse {
	return BalanceResponse{
		Balance:   MinorToMoney(s.Balance),
		Withdrawn: MinorToMoney(s.Withdrawn),
	}
}
func ToWithdrawalResponse(orders []services.ProcessedOrder) []WithdrawalOrder {
	result := make([]WithdrawalOrder, len(orders))

	for i := range orders {
		result[i] = WithdrawalOrder{
			Processed: orders[i].Processed,
			Invoice:   orders[i].Invoice,
			Amount:    MinorToMoney(orders[i].Amount),
		}
	}
	return result
}

func (s *Handler) isValidInvoice(ctx context.Context, w http.ResponseWriter, invoice string) bool {
	_, _, err := goluhn.Calculate(invoice)
	if err != nil {
		httpError := NewInvoiceError(err.Error())
		s.writeHTTPError(ctx, w, httpError)
		return false
	}
	return true
}
