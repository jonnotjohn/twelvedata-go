// Portions of this file are adapted from github.com/spacecodewor/fmpcloud-go
// Copyright (c) 2021 Igor Churbakov
// Licensed under the MIT License -- see LICENSE file for details

package twelvedata

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

type HTTPClient struct {
	logger        *zap.Logger
	client        *resty.Client
	apiKey        string
	retryCount    *int
	retryWaitTime *time.Duration
}

func (h *HTTPClient) Get(endpoint string, data map[string]string) (response *resty.Response, err error) {
	if data == nil {
		data = make(map[string]string)
	}

	data["apikey"] = h.apiKey

	retries := 0
	for retries < *h.retryCount {
		response, err = h.client.R().
			SetQueryParams(data).
			Get(endpoint)

		if err != nil || response.StatusCode() != http.StatusOK {
			time.Sleep(*h.retryWaitTime)
			retries++

			// response is not valid when there is an error
			var errOrStatusCode zapcore.Field
			if err != nil {
				errOrStatusCode = zap.Error(err)
			} else {
				errOrStatusCode = zap.Int("statusCode", response.StatusCode())
			}

			h.logger.Info(
				"Retry request",
				zap.Int("retries", retries),
				errOrStatusCode,
				zap.String("endpoint", endpoint),
				zap.Any("data", data),
			)

			continue
		}

		// If we get here, the request was successful
		break
	}

	return response, err
}
