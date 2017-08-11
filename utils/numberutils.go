package utils

import "strconv"

func ParseInt(text string) int64 {
	if len(text) == 0 {
		return 0
	}
	number, _ := strconv.ParseInt(text, 10, 64)
	return number
}

func ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}
