package utils

import "time"

func GetCurrentTimeStampSecond() int64 {
	return time.Now().UTC().Unix()
}
