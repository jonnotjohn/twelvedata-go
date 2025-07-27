package twelvedata

import (
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// TDTime handles the datetime format used by TwelveData API
type TDTime struct {
	time.Time
}

func (t *TDTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data[1 : len(data)-1])

	if parsed, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
		t.Time = parsed
		return nil
	}

	if parsed, err := time.Parse("2006-01-02", str); err == nil {
		t.Time = parsed
		return nil
	}

	// Sometimes the API returns Unix timestamps in seconds
	// This will break on 20 November 2286, but that's a long way off
	if fullStr := string(data); len(fullStr) == 10 {
		if parsedInt, errInt := strconv.ParseInt(fullStr, 10, 64); errInt == nil {
			t.Time = time.Unix(parsedInt, 0)
			return nil
		}
	}

	return errors.New("invalid datetime format encountered when parsing response from TwelveData API: " + str)
}
