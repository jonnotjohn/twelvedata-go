package twelvedata

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	urlEndpointCrypto = "/cryptocurrencies"
)

type Crypto struct {
	Symbol             string   `json:"symbol"`              // Cryptocurrency symbol (e.g. "BTC/USD", "ETH/EUR")
	AvailableExchanges []string `json:"available_exchanges"` // List of exchanges where the cryptocurrency is available (e.g. ["Binance", "Coinbase"])
	CurrencyBase       string   `json:"currency_base"`       // Base currency of the cryptocurrency (e.g. "BTC", "ETH")
	CurrencyQuote      string   `json:"currency_quote"`      // Quote currency of the cryptocurrency (e.g. "USD", "EUR")
}

type CryptoResponse struct {
	Data   []Crypto `json:"data"`   // List of cryptocurrencies
	Status string   `json:"status"` // Status of the response
}

func (c *APIClient) GetCryptocurrencies() (cryptoResponse *CryptoResponse, err error) {
	data, err := c.Client.Get(urlEndpointCrypto, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching cryptocurrencies data")
	}

	err = jsoniter.Unmarshal(data.Body(), &cryptoResponse)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling cryptocurrencies response")
	}

	return cryptoResponse, nil
}
