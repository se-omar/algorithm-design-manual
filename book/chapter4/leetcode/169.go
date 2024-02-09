package leetcode

func majorityElement(nums []int) int {
	count, curr := 0, nums[0] 

	for i := 1; i < len(nums); i++ {
		if nums[i] == curr {
			count++
		} else if count == 0 {
			curr = nums[i]
		} else {
			count--
		}
	}
	 
	return curr
}
