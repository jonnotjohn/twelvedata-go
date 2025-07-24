package objects

type TimeSeriesPeriod string
type TimeSeriesSortingOrder string

const (
	TimeSeriesPeriod1Min   TimeSeriesPeriod = "1min"
	TimeSeriesPeriod5Min   TimeSeriesPeriod = "5min"
	TimeSeriesPeriod15Min  TimeSeriesPeriod = "15min"
	TimeSeriesPeriod30Min  TimeSeriesPeriod = "30min"
	TimeSeriesPeriod45Min  TimeSeriesPeriod = "45min"
	TimeSeriesPeriod1Hour  TimeSeriesPeriod = "1hour"
	TimeSeriesPeriod2Hour  TimeSeriesPeriod = "2hour"
	TimeSeriesPeriod4Hour  TimeSeriesPeriod = "4hour"
	TimesSeriesPeriod5Hour TimeSeriesPeriod = "5hour"
	TimeSeriesPeriod1Day   TimeSeriesPeriod = "1day"
	TimeSeriesPeriod1Week  TimeSeriesPeriod = "1week"
	TimeSeriesPeriod1Month TimeSeriesPeriod = "1month"

	TimeSeriesAscending  TimeSeriesSortingOrder = "asc"
	TimeSeriesDescending TimeSeriesSortingOrder = "desc"
)

type TimeSeriesRequest struct {
	Symbol     *string
	Interval   *TimeSeriesPeriod
	OutputSize *uint16
	Order      *TimeSeriesSortingOrder
	Date       *string // ISO date format (YYYY-MM-DD)
	StartDate  *string // ISO date format with optional time (2006-01-02 or 2006-01-02 15:04:05)
	EndDate    *string // ISO date format with optional time (2006-01-02 or 2006-01-02 15:04:05)
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
	MicCode          string `json:"mic_code"` // Market Identifier Code
	Type             string `json:"type"`     // e.g., "Digital currency" or "Common stock"
}

type TimeSeriesCandle struct {
	DateTime string  `json:"datetime"` // 2025-07-23 15:30:00
	Open     float64 `json:"open"`
	Close    float64 `json:"close"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Volume   float64 `json:"volume"`
}
