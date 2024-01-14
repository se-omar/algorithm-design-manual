package chapter3

func isAnagram(s string, t string) bool {
	// populate hashmap O(n) O(n)
	if len(s) != len(t) {
		return false
	}

	tbl := map[byte]int{}

	for i := 0; i < len(s); i++ {
		tbl[s[i]]++
		tbl[t[i]]--
	}

	for _, val := range tbl {
		if val != 0 {
			return false
		}
	}

	return true
}
