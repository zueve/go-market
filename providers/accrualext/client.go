package accrualext

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/zueve/go-market/pkg/logging"
	"github.com/zueve/go-market/services/accrual"
)

type AccrualExternalClient struct {
	URL string
}

func (s *AccrualExternalClient) GetOrderStatus(ctx context.Context, order accrual.OrderVal) (Response, error) {
	url := fmt.Sprintf("%s/%s/%d", s.URL, "api/orders", order.Invoice)
	response, err := http.Get(url)
	if err != nil {
		s.log(ctx).Error().Err(err).Msg("can't parse response")
		return Response{}, err
	}
	if response.StatusCode != http.StatusOK {
		s.log(ctx).Error().Msgf("accrualExt - receive err %d", response.StatusCode)
		return Response{}, errors.New("invalid Response")
	}
	var data Response

	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		s.log(ctx).Error().Err(err).Msg("can't parse response")
		return Response{}, err
	}

	if err := json.Unmarshal(payload, &data); err != nil {
		s.log(ctx).Error().Err(err).Msg("can't parse response")
		return Response{}, err
	}

	return data, nil
}

func (s *AccrualExternalClient) log(ctx context.Context) *zerolog.Logger {
	logger := logging.GetLogger(ctx).With().
		Str(logging.Source, "accrual.Client").
		Str(logging.Layer, "providers").
		Logger()
	return &logger
}
