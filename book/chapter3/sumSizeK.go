package main

func sumSizeK(nums []int, k int) int {
	// naive approach

	total := 0
	arrSum := 0

	for i := 0; i < len(nums)-k; i++ {
		for j := i; j < k; j++ {
			arrSum += nums[j]
		}
		total += arrSum
	}

	return arrSum
}
