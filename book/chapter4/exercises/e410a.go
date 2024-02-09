package exercises

import "golang.org/x/exp/slices"

func twoSumA(s []int, x int) bool{
	slices.Sort(s)
	l, r := 0, len(s)-1

	for l < r {
		if s[l]+s[r] < x {
			l++
		} else if s[l]+s[r] > x {
			r--
		} else {
			return true
		}
	}

	return false
}
