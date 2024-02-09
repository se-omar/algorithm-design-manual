package leetcode

func majorityElement(nums []int) int {
	half := len(nums) / 2
	hm := make(map[int]int)
	for _, v := range nums {
		hm[v]++
	}

	for k, v := range hm {
		if v > half {
			return k
		}
	}

	return 0

}
