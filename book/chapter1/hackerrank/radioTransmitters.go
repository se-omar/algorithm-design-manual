package hackerrank

import (
	"sort"
)

func HackerlandRadioTransmitters(x []int32, k int32) int32 {
	arr := convertSlice(x)
	rng := int(k)

	sort.Ints(arr)
	res := int32(0)

	for len(arr) > 0 {
		start := arr[0]
		center := start + rng

		for indexOf(arr, center) == -1 && center < arr[len(arr)-1] {
			center -= 1
		}

		res += 1
		end := center + rng  

		for len(arr) > 0 && arr[0] <= end {
			arr = arr[1:]
		}
	}

	return res
}

func indexOf(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func convertSlice(x []int32) []int {
	var res []int
	for _, v := range x {
		res = append(res, int(v))
	}
	return res
}
