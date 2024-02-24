package time

import (
	"fmt"
	"time"
)

var location, _ = time.LoadLocation("Asia/Jakarta")

func GetCurrentTime() time.Time {
	return time.Now().UTC().In(location)
}

func StrToLocalTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s+"Z")
	if err != nil {
		return time.Time{}, fmt.Errorf("error while parsing string to local time")
	}
	t = t.Add(-7 * time.Hour).In(location)
	return t, nil
}
