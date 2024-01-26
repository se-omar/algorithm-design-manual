package main

func stringInMagzine(s, m string) bool {
	var ht map[rune]int

	for _, v := range s {
		ht[v]++
	}

	for _, v := range m {
		if _, ok := ht[v]; ok && ht[v] > 0 {
			ht[v]--
		} else {
			delete(ht, v) 
		}
	}

		if len(ht) > 0 {
			return false
		}

	return true
}
