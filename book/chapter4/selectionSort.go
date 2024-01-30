package chapter4

func SelectionSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		minIdx := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minIdx] {
				minIdx = j
			}
		}

		s[i], s[minIdx] = s[minIdx], s[i]
	}

	return s
}
