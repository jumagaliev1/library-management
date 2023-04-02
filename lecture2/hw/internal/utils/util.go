package utils

import (
	"fmt"
	"strconv"
)

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
		return -1
	}
	return n
}
