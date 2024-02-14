package leetcode

import "golang.org/x/exp/slices"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	prevEnd := intervals[0][1]
	count := 0

	for i := 1; i < len(intervals); i++ {
		currStart := intervals[i][0]
		currEnd := intervals[i][1]
		if currStart < prevEnd {
			count++
			prevEnd = min(prevEnd, currEnd)
		} else {
			prevEnd = currEnd
		}
	}

	return count
}
