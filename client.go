package twelvedata

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type APIUrl string

type Config struct {
	Logger        *zap.Logger
	HTTPClient    *resty.Client
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
}
