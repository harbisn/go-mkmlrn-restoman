package time

import "time"

var location, _ = time.LoadLocation("Asia/Jakarta")

func GetCurrentTime() time.Time {
	return time.Now().UTC().In(location)
}
