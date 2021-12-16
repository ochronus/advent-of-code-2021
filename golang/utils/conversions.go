package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func StrToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Hex2Bin(in string) (binaryStr string) {
	bytes, _ := hex.DecodeString(in)
	for _, b := range bytes {
		binaryStr += fmt.Sprintf("%08b", b)
	}
	return
}

func Bin2Dec(bin string) int64 {
	num, _ := strconv.ParseInt(bin, 2, 64)
	return num
}
