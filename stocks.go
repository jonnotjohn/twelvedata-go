package twelvedata

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	urlEndpointStocks = "/stocks"
)

type Stocks struct {
	Symbol   string `json:"symbol"`    // Stock symbol (e.g. "AAPL")
	Name     string `json:"name"`      // Full name of the stock (e.g. "Apple Inc.")
	Currency string `json:"currency"`  // Currency code (e.g. "USD")
	Exchange string `json:"exchange"`  // Exchange code (e.g. "NASDAQ")
	MicCode  string `json:"mic_code"`  // Market Identifier Code (e.g. "XNAS" for NASDAQ)
	Country  string `json:"country"`   // Country code (e.g. "US")
	Type     string `json:"type"`      // Type of the stock (e.g. "Common Stock", "ETF")
	FigiCode string `json:"figi_code"` // Financial Instrument Global Identifier
	CfiCode  string `json:"cfi_code"`  // Classification of Financial Instruments code
	ISIN     string `json:"isin"`      // International Securities Identification Number
	CUSIP    string `json:"cusip"`     // Committee on Uniform Securities Identification Procedures
}

type StocksResponse struct {
	Data   []Stocks `json:"data"`  // List of stocks
	Count  int      `json:"count"` // Total number of stocks
	Status string   `json:"status"`
}

func (c *APIClient) GetStocks() (stocksResponse *StocksResponse, err error) {
	data, err := c.Client.Get(urlEndpointStocks, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching stocks data")
	}

	err = jsoniter.Unmarshal(data.Body(), &stocksResponse)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling stocks response")
	}

	return stocksResponse, nil
}
