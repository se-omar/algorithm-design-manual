package exercises

func findCount(s []int) int {
	hm := make(map[int]int)
	half := len(s) / 2
	for _, v := range s {
		hm[v]++
	}

	for k, v := range hm {
		if v >= half {
			return k
		}
	}

	return 0
}
