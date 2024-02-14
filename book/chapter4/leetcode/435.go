package leetcode

import "golang.org/x/exp/slices"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	output := []int{intervals[0][0], intervals[0][1]}
	count := 0

	for i := 1; i < len(intervals); i++ {
		lastEnd := output[1]
		currStart := intervals[i][0]
		currEnd := intervals[i][1]
		if currStart < lastEnd {
			count++
			newOp := []int{}
			if lastEnd < currEnd {
				newOp = output
			} else {
				newOp = []int{currStart, currEnd}
			}
			output = newOp
		} else {
			output = []int{currStart, currEnd}
		}
	}

	return count
}
