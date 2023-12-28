// https://leetcode.com/problems/counting-bits

package chapter2

func CountBits(n int) []int {
	arr := make([]int, n+1)
	for i := n; i > 0; i-- {
		v := i
		for v > 0 {
			if v%2 == 1 {
				arr[i]++
			}
			v = v / 2
		}

	}

	return arr
}
