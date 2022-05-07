package dates

import (
	"time"
)

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetCurrentTime() string {
	DateCreated := GetNow().Format(apiDateLayout)
	return DateCreated
}
