package main

func SumSizeK(nums []int, k int) []int {
	// sliding window
	total := []int{}

	arrSum := sum(nums[:k])
	total = append(total, arrSum)

	for i := 1; i <= len(nums)-k; i++ {
		arrSum = arrSum - nums[i-1] + nums[i+k-1]
		total = append(total, arrSum)
	}

	return total
}

func sum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

