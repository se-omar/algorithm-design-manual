package hackerrank

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2 > v1 || v2 == v1 {
		return "NO"
	}

	if (x1-x2)%(v2-v1) != 0 {
		return "NO"
	}

	return "YES"
}
