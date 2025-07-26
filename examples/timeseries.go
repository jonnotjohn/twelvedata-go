package main

import (
	"fmt"
	"github.com/jonnotjohn/twelvedata-go"
)

func main() {
	apiClient, err := twelvedata.NewAPIClient(twelvedata.Config{
		APIKey: "demo", // Replace with your actual API key
	})
	if err != nil {
		panic("Error creating API client: " + err.Error())
	}

	symbol := "AAPL"
	interval := twelvedata.TimeSeriesInterval15Min

	tsReq := twelvedata.TimeSeriesRequest{
		Symbol:   &symbol,
		Interval: &interval,
	}

	candles, err := apiClient.GetTimeSeries(tsReq)

	if err != nil {
		panic(fmt.Sprintf("Error fetching time series data: %v", err))
	}
	fmt.Printf("Time Series Data for %s: %+v\n", symbol, candles)
}
