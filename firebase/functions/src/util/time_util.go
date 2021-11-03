package util

import (
	"math"
	"strconv"
	"time"
)

func ParseTimeFromFloatStr(floatStr string) (*time.Time, error) {
	float, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		return nil, err
	}

	t := ParseTimeFromFloat(float)
	return &t, nil
}

func ParseTimeFromFloat(float float64) time.Time {
	sec, dec := math.Modf(float)
	return time.Unix(int64(sec), int64(dec*(1e9)))
}
