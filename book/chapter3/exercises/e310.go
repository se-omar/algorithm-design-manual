package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	// sort both strings O(nlogn) O(1)
	if len(s) != len(t) {
		return false
	}

	x := strings.Split(s, "")
	sort.Strings(x)
	s = strings.Join(x, "")

	x = strings.Split(t, "")
	sort.Strings(x)
	t = strings.Join(x, "")

	return s == t
}
