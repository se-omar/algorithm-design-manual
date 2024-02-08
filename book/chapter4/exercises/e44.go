package exercises

type pair struct {
	num   int
	color string
}

func sortColorPairs(pairs []pair) []pair {
	reds, blues, yellows := []pair{}, []pair{}, []pair{}

	for _, v := range pairs {
		if v.color == "red" {
			reds = append(reds, v)
		} else if v.color == "blues" {
			blues = append(blues, v)
		} else {
			yellows = append(yellows, v)
		}
	}

	combined := []pair{}
	combined = append(combined, reds...)
	combined = append(combined, blues...)
	combined = append(combined, yellows...)

	return combined
}
