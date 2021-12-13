package utils

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func StringIntersection(a, b string) (count int) {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			count++
		}
	}
	return
}

func StrToIntList(s string, separator string) []int {
	var list []int
	for _, word := range strings.Split(s, separator) {
		list = append(list, StrToInt(word))
	}
	return list
}
