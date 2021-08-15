package utils

import (
	"log"
	"math"
	"strconv"
	"time"
)

func ParseTime(floatStr string) (t time.Time, err error) {
	float, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		log.Fatalf("ParseFoat failed.(err=%+v)", err)
		return t, err
	}

	sec, dec := math.Modf(float)
	t = time.Unix(int64(sec), int64(dec*(1e9)))

	return t, nil
}
