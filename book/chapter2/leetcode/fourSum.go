// https://leetcode.com/problems/4sum/

package chapter2

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	currQuad := []int{}
	sort.Ints(nums)

	var kSum func(k int, start int, target int)
	kSum = func(k int, start int, target int) {
		if k != 2 {
			for i := start; i < len(nums)-k + 1; i++ {
				if i > start && nums[i] == nums[i-1] {
					continue
				}
				currQuad = append(currQuad, nums[i])
				kSum(k-1, i+1, target-nums[i])
				currQuad = currQuad[:len(currQuad)-1]

			}

		} else {
			l, r := start, len(nums)-1
			for l < r {
				sum := nums[l] + nums[r]

				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					quad := append(currQuad, nums[l], nums[r])
					ans = append(ans, quad)
					l++
					for nums[l] == nums[l-1] && l < r {
						l++
					}
				}

			}
		}

	}

	kSum(4, 0, target)

	return ans
}
