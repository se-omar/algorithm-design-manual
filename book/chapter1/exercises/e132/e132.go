// Write a function to perform integer division without using either the / or *
// operators. Find a fast way to do it

package e132

func solution(numer, denom int) int {
	if denom > numer {
		return 0
	}

	if denom == numer {
		return 1
	}

	res := 1
	for {
		numer -= denom
		if numer <= 0 {
			break
		}
		res += 1
	}

	return res
}
