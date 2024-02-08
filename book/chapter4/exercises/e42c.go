package exercises

import (
	"golang.org/x/exp/slices"
	"math"
)

func MinDiff(s []int) (int, int) {
	slices.Sort(s)
	num1, num2 := s[0], s[0]
	mini := math.MaxInt
	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}

		currMini := math.Abs(float64(s[i]-s[i-1]))
		if currMini < float64(mini) {
			mini = int(currMini)
			num1 = s[i]
			num2 = s[i-1]
		}
	}

	return num1, num2
}
