// implements the nearestNeighbor Huerestic for
// the robot tour optimization problem

package e130

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func solution(P []Point) []Point {
	p0 := P[0]
	p := p0
	i := 0
	visited := []Point{}
	setLen := len(P)

	for i < setLen {
		i += 1
		p1, idx := getNearestPoint(p, P)
		visited = append(visited, p1)
		P = remove(P, idx)
		p = p1
	}
	visited = append(visited, p0)
	return visited
}

func getDist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func getNearestPoint(p Point, S []Point) (Point, int) {
	minDist := getDist(p, S[0])
	idx := 0
	pnt := S[0]
	i := 0
	for i < len(S) {
		dist := getDist(p, S[i])
		if dist < minDist {
			minDist = dist
			idx = i
			pnt = S[i]
		}
		i += 1
	}
	return pnt, idx
}

func remove(s []Point, i int) []Point {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
