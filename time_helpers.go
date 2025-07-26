package twelvedata

import (
	"time"
)

// TDTime handles the datetime format used by TwelveData API
type TDTime struct {
	time.Time
}

func (t *TDTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data[1 : len(data)-1])

	parsed, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}
