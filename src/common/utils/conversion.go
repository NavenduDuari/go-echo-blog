package utils

import "strconv"

func Int64ToStr(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
