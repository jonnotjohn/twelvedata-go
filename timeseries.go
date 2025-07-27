package twelvedata

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"time"
)

const (
	UrlEndpointTimeSeries = "/time_series"
)

type TimeSeriesInterval string

const (
	TimeSeriesInterval1Min   TimeSeriesInterval = "1min"
	TimeSeriesInterval5Min   TimeSeriesInterval = "5min"
	TimeSeriesInterval15Min  TimeSeriesInterval = "15min"
	TimeSeriesInterval30Min  TimeSeriesInterval = "30min"
	TimeSeriesInterval45Min  TimeSeriesInterval = "45min"
	TimeSeriesInterval1Hour  TimeSeriesInterval = "1hour"
	TimeSeriesInterval2Hour  TimeSeriesInterval = "2hour"
	TimeSeriesInterval4Hour  TimeSeriesInterval = "4hour"
	TimeSeriesInterval5Hour  TimeSeriesInterval = "5hour"
	TimeSeriesInterval1Day   TimeSeriesInterval = "1day"
	TimeSeriesInterval1Week  TimeSeriesInterval = "1week"
	TimeSeriesInterval1Month TimeSeriesInterval = "1month"
)

// TimeSeriesRequest is the available parameters for a time series request
type TimeSeriesRequest struct {
	Symbol        *string             // Required: Symbol of the asset (e.g. "AAPL", "BTC/USD")
	FIGI          *string             // Financial Instrument Global Identifier
	ISIN          *string             // International Securities Identification Number
	CUSIP         *string             // Committee on Uniform Securities Identification Procedures
	Interval      *TimeSeriesInterval // Required: Time interval for the candles (e.g. "1min", "1day")
	Exchange      *string             // Exchange code (e.g. "NASDAQ", "Binance")
	MicCode       *string             // Market Identifier Code (e.g. "XNAS" for NASDAQ)
	Country       *string             // Country code (e.g. "US" or "United States")
	Type          *string             // Type of asset (e.g. "Digital currency", "Common stock")
	OutputSize    *int                // Number of candles to return (default is 30, max is 5000)
	PrePost       *bool               // Include pre/post market data (default is false)
	DP            *int                // Number of decimal places for float values. Supports 0-11, default is -1 (API automatically determines precision)
	Order         *string             // Sorting order for the results "asc" and "desc" (default is "desc")
	TimeZone      *string             // Timezone for the response (e.g. "America/New_York", "UTC"). Defaults to "Exchange"
	Date          *time.Time          // Specific day to fetch data for (time is ignored)
	StartDate     *time.Time          // Time when the series starts
	EndDate       *time.Time          // Time when the series ends
	PreviousClose *bool               // Include previous close price in the response (default is false)
	Adjust        *string             // Adjusting mode for prices ("none", "dividends", "splits", "all"). Default is "none"
}

func (req TimeSeriesRequest) ToParams() (map[string]string, error) {
	params := make(map[string]string)

	if req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}

	if req.Interval == nil {
		return nil, errors.New("interval is required")
	}
	params["interval"] = string(*req.Interval)

	AddStringParam(params, "symbol", req.Symbol)
	AddStringParam(params, "figi", req.FIGI)
	AddStringParam(params, "isin", req.ISIN)
	AddStringParam(params, "cusip", req.CUSIP)
	AddStringParam(params, "exchange", req.Exchange)
	AddStringParam(params, "mic_code", req.MicCode)
	AddStringParam(params, "country", req.Country)
	AddStringParam(params, "type", req.Type)
	AddStringParam(params, "timezone", req.TimeZone)
	AddStringParam(params, "adjust", req.Adjust)
	AddStringParam(params, "order", req.Order)

	AddIntParam(params, "outputsize", req.OutputSize)
	AddIntParam(params, "dp", req.DP)

	AddBoolParam(params, "prepost", req.PrePost)
	AddBoolParam(params, "previous_close", req.PreviousClose)

	AddDateParam(params, "date", req.Date, "2006-01-02")
	AddDateParam(params, "start_date", req.StartDate, "2006-01-02 15:04:05")
	AddDateParam(params, "end_date", req.EndDate, "2006-01-02 15:04:05")

	return params, nil
}

type TimeSeriesResponse struct {
	Meta    TimeSeriesResponseMeta `json:"meta"`
	Candles []TimeSeriesCandle     `json:"values"`
}

type TimeSeriesResponseMeta struct {
	Symbol           string `json:"symbol"`
	Interval         string `json:"interval"`
	Currency         string `json:"currency"`
	CurrencyBase     string `json:"currency_base"`
	CurrencyQuote    string `json:"currency_quote"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	Type             string `json:"type"`
}

type TimeSeriesCandle struct {
	DateTime TDTime  `json:"datetime"`
	Open     float64 `json:"open,string"`
	Close    float64 `json:"close,string"`
	High     float64 `json:"high,string"`
	Low      float64 `json:"low,string"`
	Volume   float64 `json:"volume,string"`
}

func (c *APIClient) GetTimeSeries(req TimeSeriesRequest) (candles *TimeSeriesResponse, err error) {
	params, err := req.ToParams()
	if err != nil {
		return nil, errors.Wrap(err, "Error converting TimeSeriesRequest to params")
	}

	resp, err := c.Client.Get(UrlEndpointTimeSeries, params)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching time series data")
	}

	err = jsoniter.Unmarshal(resp.Body(), &candles)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling time series response")
	}

	return candles, nil
}
