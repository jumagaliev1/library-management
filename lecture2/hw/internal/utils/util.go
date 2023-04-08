package utils

import (
	"strconv"
)
// в го не используется папка utils
func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return n
}
