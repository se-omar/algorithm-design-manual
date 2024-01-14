package chapter3

func isAnagram(s string, t string) bool {
	// populate array O(n) O(n)
	if len(s) != len(t) {
		return false
	}

	var store [128]int
	for i := 0; i < len(s); i++ {
		store[int(s[i])]++
		store[int(t[i])]--
	}

	for _, val := range store {
		if val != 0 {
			return false
		}

	}

	return true

}
