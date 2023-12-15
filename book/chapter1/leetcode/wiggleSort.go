package leetcode

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	idx := len(sorted) - 1
	for i := 1; i < len(nums); i += 2 {
		nums[i] = sorted[idx]
		idx -= 1
	}

	for i := 0; i < len(nums); i += 2 {
		nums[i] = sorted[idx]
		idx -= 1
	}


	return nums

}
