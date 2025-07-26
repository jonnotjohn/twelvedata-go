package twelvedata

import (
	"strconv"
	"time"
)

func AddStringParam(params map[string]string, key string, value *string) {
	if value != nil {
		params[key] = *value
	}
}

func AddIntParam(params map[string]string, key string, value *int) {
	if value != nil {
		params[key] = strconv.Itoa(*value)
	}
}

func AddBoolParam(params map[string]string, key string, value *bool) {
	if value != nil {
		params[key] = strconv.FormatBool(*value)
	}
}

func AddDateParam(params map[string]string, key string, value *time.Time, format string) {
	if value != nil {
		params[key] = value.Format(format)
	}
}
