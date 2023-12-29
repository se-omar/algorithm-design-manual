// https://leetcode.com/problems/two-sum

package chapter2

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		el := nums[i]
		if _, ok := mp[target-el]; ok {
			return []int{i, mp[target-el]}
		}

		mp[el] = i
	}

	return []int{}
}
