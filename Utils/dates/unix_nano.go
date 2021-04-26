package dates

import (
	"time"
)

func EpchoConvertor() int64 {
	now := time.Now()
	secs := now.Unix()
	return secs
}
