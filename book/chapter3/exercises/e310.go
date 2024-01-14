package chapter3

import "sort"

func isAnagram(s string, t string) bool {
	// sort both strings O(nlogn) O(1)
	if len(s) != len(t) {
		return false
	}

	bytesS := []byte(s)
	bytesT := []byte(t)
	sort.Slice(bytesS, func(i, j int) bool { return bytesS[i] < bytesS[j] })
	sort.Slice(bytesT, func(i, j int) bool { return bytesT[i] < bytesT[j] })
	s = string(bytesS)
	t = string(bytesT)


	return s == t
}
