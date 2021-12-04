package utils

import "strconv"

func StrToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
