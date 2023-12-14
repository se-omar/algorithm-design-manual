package e130

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	points := []Point{
		{x: -5, y: 0},
		{x: 0, y: 0},
		{x: 1, y: 0},
		{x: 11, y: 0},
		{x: -1, y: 0},
		{x: -21, y: 0},
		{x: 3, y: 0},
	}

	correctOrder := []Point{
		{x: -5, y: 0},
		{x: -1, y: 0},
		{x: 0, y: 0},
		{x: 1, y: 0},
		{x: 3, y: 0},
		{x: 11, y: 0},
		{x: -21, y: 0},
		{x: -5, y: 0},
	}

	assert.Equal(t, solution(points), correctOrder, "wrong ordering")
}
