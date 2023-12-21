package chapter1

func IndexOf(t, s string) int {
	idx := -1
	for i := 0; i < len(t); i++ {
		idx = i
		for j := 0; j < len(s); j++ {
			if s[j] != t[i+j] {
				idx = -1
				break
			}
		}

		if idx != -1 {
			break
		}

	}

	return idx
}
