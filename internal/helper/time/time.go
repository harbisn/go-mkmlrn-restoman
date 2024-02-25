package time

import (
	"fmt"
	"time"
)

func GetLocation() *time.Location {
	location, _ := time.LoadLocation("Asia/Jakarta")
	return location
}

func GetCurrentTime() time.Time {
	return time.Now().UTC().In(GetLocation()).Truncate(time.Millisecond)
}

func StrToLocalTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s+"Z")
	if err != nil {
		return time.Time{}, fmt.Errorf("error while parsing string to local time")
	}
	t = t.Add(-7 * time.Hour).In(GetLocation())
	return t, nil
}
