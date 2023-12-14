// https://leetcode.com/problems/next-greater-element-i
package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	mp := make(map[int]int)
	var stk []int

	for idx, value := range nums1 {
		mp[value] = idx
	}

	for _, value := range nums2 {
		for len(stk) > 0 && value > stk[len(stk)-1] {
			el := stk[len(stk)-1]
			idx := mp[el]
			stk = stk[:len(stk)-1]
			ans[idx] = value
		}
		if _, ok := mp[value]; ok {
			stk = append(stk, value)
		}
	}

	return ans
}
