// https://leetcode.com/problems/daily-temperatures

package leetcode

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	var stk []int

	for i, value := range temperatures {
		for len(stk) > 0 && value > temperatures[stk[len(stk)-1]] {
			idx := stk[len(stk)-1] 
			stk = stk[:len(stk)-1]
			answer[idx] = i - idx
		}

		stk = append(stk, i)
	}

	return answer
}
