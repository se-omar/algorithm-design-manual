package chapter1

func matrixMultiply(x [][]int, y [][]int) [][]int {
	z := [][]int{}

	for i := 0; i < len(x[0]); i++ {
		for j := 0; j < len(y[1]); j++ {
			z[i][j] = 0
			for k := 0; k < len(x[0]); k++ {
				z[i][j] += x[i][k] * y[k][j]
			}

		}
	}

	return z
}
