package twelvedata

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	UrlEndpointQuote = "/quote"
)

type QuoteInterval string

const (
	QuoteInterval1Min   QuoteInterval = "1min"
	QuoteInterval5Min   QuoteInterval = "5min"
	QuoteInterval15Min  QuoteInterval = "15min"
	QuoteInterval30Min  QuoteInterval = "30min"
	QuoteInterval45Min  QuoteInterval = "45min"
	QuoteInterval1Hour  QuoteInterval = "1h"
	QuoteInterval2Hour  QuoteInterval = "2h"
	QuoteInterval4Hour  QuoteInterval = "4h"
	QuoteInterval1Day   QuoteInterval = "1day"
	QuoteInterval1Week  QuoteInterval = "1week"
	QuoteInterval1Month QuoteInterval = "1month"
)

// QuoteRequest is the available parameters for a quote request
type QuoteRequest struct {
	Symbol           *string        // Required: Symbol of the asset (e.g. "AAPL", "BTC/USD")
	FIGI             *string        // Financial Instrument Global Identifier
	ISIN             *string        // International Securities Identification Number
	CUSIP            *string        // Committee on Uniform Securities Identification Procedures
	Interval         *QuoteInterval // Time interval for the quotes (e.g. "1min", "1week") defaults to "1day"
	Exchange         *string        // Exchange code (e.g. "NASDAQ", "Binance")
	MicCode          *string        // Market Identifier Code (e.g. "XNAS" for NASDAQ)
	Country          *string        // Country code (e.g. "US" or "United States")
	VolumeTimePeriod *int           // Number of periods for Average Volume
	Type             *string        // Type of asset (e.g. "Digital currency", "Common stock")
	PrePost          *bool          // Include pre/post market data (default is false)
	EOD              *bool          // If true, then return data for closed day
	RollingPeriod    *int           // Number of hours for calculate rolling change at period
	DP               *int           // Number of decimal places for float values. Supports 0-11, default is 5
	TimeZone         *string        // Timezone for the response (e.g. "America/New_York", "UTC"). Defaults to "Exchange"
}

func (req QuoteRequest) ToParams() (map[string]string, error) {
	params := make(map[string]string)

	if req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}

	if req.Interval != nil {
		params["interval"] = string(*req.Interval)
	}

	AddStringParam(params, "symbol", req.Symbol)
	AddStringParam(params, "figi", req.FIGI)
	AddStringParam(params, "isin", req.ISIN)
	AddStringParam(params, "cusip", req.CUSIP)
	AddStringParam(params, "exchange", req.Exchange)
	AddStringParam(params, "mic_code", req.MicCode)
	AddStringParam(params, "country", req.Country)
	AddStringParam(params, "type", req.Type)
	AddStringParam(params, "timezone", req.TimeZone)

	AddIntParam(params, "volume_time_period", req.VolumeTimePeriod)
	AddIntParam(params, "rolling_period", req.RollingPeriod)
	AddIntParam(params, "dp", req.DP)

	AddBoolParam(params, "prepost", req.PrePost)
	AddBoolParam(params, "eod", req.EOD)

	return params, nil
}

type QuoteResponse struct {
	Symbol                string            `json:"symbol"`
	Name                  string            `json:"name"`
	Exchange              string            `json:"exchange"`
	MicCode               string            `json:"mic_code"`
	Currency              string            `json:"currency"`
	DateTime              TDTime            `json:"datetime"`
	Timestamp             TDTime            `json:"timestamp"`
	LastQuoteAt           TDTime            `json:"last_quote_at"`
	Open                  float64           `json:"open,string"`
	High                  float64           `json:"high,string"`
	Low                   float64           `json:"low,string"`
	Close                 float64           `json:"close,string"`
	Volume                float64           `json:"volume,string"`
	PreviousClose         float64           `json:"previous_close,string"`
	Change                float64           `json:"change,string"`
	PercentChange         float64           `json:"percent_change,string"`
	AverageVolume         float64           `json:"average_volume,string"`
	Rolling1DayChange     float64           `json:"rolling_1day_change,string"`
	Rolling7DayChange     float64           `json:"rolling_7day_change,string"`
	RollingPeriodChange   float64           `json:"rolling_period_change,string"`
	IsMarketOpen          bool              `json:"is_market_open"`
	FiftyTwoWeek          QuoteFiftyTwoWeek `json:"fifty_two_week"`
	ExtendedChange        float64           `json:"extended_change,string"`
	ExtendedPercentChange float64           `json:"extended_percent_change,string"`
	ExtendedPrice         float64           `json:"extended_price,string"`
	ExtendedTimestamp     int64             `json:"extended_timestamp"`
}

type QuoteFiftyTwoWeek struct {
	Low               float64 `json:"low,string"`
	High              float64 `json:"high,string"`
	LowChange         float64 `json:"low_change,string"`
	HighChange        float64 `json:"high_change,string"`
	LowChangePercent  float64 `json:"low_change_percent,string"`
	HighChangePercent float64 `json:"high_change_percent,string"`
	Range             string  `json:"range"`
}

func (c *APIClient) GetQuote(req QuoteRequest) (quote *QuoteResponse, err error) {
	params, err := req.ToParams()
	if err != nil {
		return nil, errors.Wrap(err, "Error converting QuoteRequest to params")
	}

	data, err := c.Client.Get(UrlEndpointQuote, params)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching quote data")
	}

	err = jsoniter.Unmarshal(data.Body(), &quote)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling quote response")
	}

	return quote, nil
}
