package leetcode

import "golang.org/x/exp/slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	output := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastEnd := output[len(output)-1][1]
		currStart := intervals[i][0]
		currEnd := intervals[i][1]
		if currStart <= lastEnd {
			output[len(output)-1][1] = max(currEnd, lastEnd)
		} else {
			output = append(output, intervals[i])
		}
	}

	return output
}
