package twelvedata

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	urlEndpointLogo = "/logo"
)

type LogoRequest struct {
	Symbol   *string // Required: Symbol of the asset (e.g. "AAPL", "BTC/USD")
	Exchange *string // Exchange code (e.g. "NASDAQ", "Binance")
	MicCode  *string // Market Identifier Code (e.g. "XNAS" for NASDAQ)
	Country  *string // Country code (e.g. "US" or "United States")
}

func (req LogoRequest) ToParams() (map[string]string, error) {
	params := make(map[string]string)

	if req.Symbol == nil {
		return nil, errors.New("symbol is required")
	}

	AddStringParam(params, "symbol", req.Symbol)
	AddStringParam(params, "exchange", req.Exchange)
	AddStringParam(params, "mic_code", req.MicCode)
	AddStringParam(params, "country", req.Country)

	return params, nil
}

type LogoMeta struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
}

type Logo struct {
	Meta      LogoMeta `json:"meta"`
	URL       string   `json:"url"`        // URL of the logo image (for stocks only)
	LogoBase  string   `json:"logo_base"`  // URL of the logo of the base currency (for crypto and forex only)
	LogoQuote string   `json:"logo_quote"` // URL of the logo of the quote currency (for crypto and forex only)
}

func (c *APIClient) GetLogo(req LogoRequest) (logo *Logo, err error) {
	params, err := req.ToParams()
	if err != nil {
		return nil, errors.Wrap(err, "Error converting LogoRequest to params")
	}

	data, err := c.Client.Get(urlEndpointLogo, params)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching logo data")
	}

	err = jsoniter.Unmarshal(data.Body(), &logo)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling logo response")
	}

	return logo, nil
}
