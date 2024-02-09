package exercises

func ListMode(s []int) int {
	hm := make(map[int]int)
	for _, v := range s {
		hm[v]++
	}

	maxi, key := 0, 0
	for k, v := range hm {
		if v > maxi {
			maxi = v
			key = k
		}
	}

	return key
}
