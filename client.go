// Portions of this file are adapted from github.com/spacecodewor/fmpcloud-go
// Copyright (c) 2021 Igor Churbakov
// Licensed under the MIT License -- see LICENSE file for details

package twelvedata

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type APIUrl string

const (
	apiTwelveDataURL  APIUrl = "https://api.twelvedata.com"
	apiKeyDefault            = "demo"
	apiTimeoutDefault        = 25
)

type Config struct {
	Logger        *zap.Logger
	RestyClient   *resty.Client
	APIKey        string
	APIUrl        APIUrl
	Debug         bool
	RetryCount    *int
	RetryWaitTime *time.Duration
	Timeout       int
}

type APIClient struct {
	Logger *zap.Logger
	Debug  bool
	Client *HTTPClient
}

// NewAPIClient creates a new API client
func NewAPIClient(cfg Config) (*APIClient, error) {
	APIClient := &APIClient{Logger: cfg.Logger, Debug: cfg.Debug}
	if APIClient.Logger == nil {
		logger, err := createNewLogger()
		if err != nil {
			return nil, errors.Wrap(err, "Error creating new zap logger")
		}

		APIClient.Logger = logger
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = apiTimeoutDefault
	}

	// Init resty client
	if cfg.RestyClient == nil {
		cfg.RestyClient = resty.New()
	}

	cfg.RestyClient.SetDebug(APIClient.Debug)
	cfg.RestyClient.SetTimeout(time.Duration(cfg.Timeout) * time.Second)

	if len(cfg.APIUrl) == 0 {
		cfg.APIUrl = apiTwelveDataURL
	}

	if len(cfg.APIKey) == 0 {
		cfg.APIKey = apiKeyDefault
	}

	cfg.RestyClient.SetBaseURL(string(cfg.APIUrl))

	HTTPClient := &HTTPClient{
		client: cfg.RestyClient,
		apiKey: cfg.APIKey,
		logger: APIClient.Logger,
	}

	if cfg.RetryCount != nil {
		HTTPClient.retryCount = cfg.RetryCount
	}

	if cfg.RetryWaitTime != nil {
		HTTPClient.retryWaitTime = cfg.RetryWaitTime
	}

	if HTTPClient.retryCount == nil || *HTTPClient.retryCount == 0 {
		retryCount := 1
		HTTPClient.retryCount = &retryCount
	}

	if HTTPClient.retryWaitTime == nil || *HTTPClient.retryWaitTime == 0 {
		retryWaitTime := 1 * time.Second
		HTTPClient.retryWaitTime = &retryWaitTime
	}

	APIClient.Client = HTTPClient

	return APIClient, nil
}

func createNewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Logger Error: init")
	}

	return logger, nil
}
