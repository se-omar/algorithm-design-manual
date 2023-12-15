package hackerrank

func rotateLeft(d int32, arr []int32) []int32 {
	x := int(d) % len(arr)

	if x == len(arr) || len(arr) == 0 {
		return arr
	}

	for i := 0; i < x; i++ {
		first := arr[0]
		arr = arr[1:]
		arr = append(arr, first)
	}

	return arr
}
