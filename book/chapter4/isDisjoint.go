package chapter4

func IsDisjoint(s1, s2 []int) bool {
	// naive solution
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				return true
			}
		}
	}
	return false

}
