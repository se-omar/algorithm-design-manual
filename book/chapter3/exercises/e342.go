package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	sarr := strings.Split(s, " ")
	s2 := []string{}
	for i := len(sarr) - 1; i >= 0; i-- {
		if len(sarr[i]) > 0 && !strings.Contains(sarr[i], " ") {
			s2 = append(s2, sarr[i])
		}
	}

	return strings.Join(s2, " ")
}
