// An unsorted array of size n contains distinct integers between 1 and n + 1,
// with one element missing. Give an O(n) algorithm to find the missing integer,
// without using any extra space

package chapter2

func solution(arr []int) int {
	arrLen := len(arr)
	actualSum := ((arrLen + 1) * (arrLen + 2)) / 2
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return actualSum - sum
}
