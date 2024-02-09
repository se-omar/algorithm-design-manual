package exercises

import "golang.org/x/exp/slices"

func FindPair(s1, s2 []int, x int) bool {
	slices.Sort(s1)
	slices.Sort(s2)

	s1p := 0
	s2p := len(s2) - 1

	for i := 0; i < len(s1); i++ {
		if s1[s1p]+s2[s2p] < x {
			s1p++
		} else if s1[s1p]+s2[s2p] > x {
			s2p--
		} else {
			return true
		}
	}

	return false
}
