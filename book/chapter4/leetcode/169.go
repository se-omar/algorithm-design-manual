package leetcode

func majorityElement(nums []int) int {
	count, curr := 0, 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			curr = nums[i]
		}

		if nums[i] == curr {
			count++
		} else {
			count--
		}
	}

	return curr
}
