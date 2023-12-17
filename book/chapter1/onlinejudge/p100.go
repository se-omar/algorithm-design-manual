// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=29&page=show_problem&problem=36
package onlinejudge


func maximumCycle(n1, n2 int) int {
	maxCnt := 0

	for i := n1; i < n2; i++ {
		n3 := i
		cnt := 1
		for n3 != 1 {
			if n3%2 == 1 {
				n3 = 3*n3 + 1
			} else {
				n3 = n3 / 2
			}
			cnt += 1
		}

		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	return maxCnt
}
