package exercises

func MaxDiff(s []int) (int, int) {
	maxi, mini := s[0], s[0]
	for _, v := range s {
		if v > maxi {
			maxi = v
		}

		if v < mini {
			mini = v
		}
	}

	return maxi, mini
}
