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
