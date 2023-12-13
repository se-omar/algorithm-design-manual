package e132

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 5, solution(10, 2), "invalid value")
	assert.Equal(t, 5, solution(15, 3), "invalid value")
	assert.Equal(t, 50, solution(1000, 20), "invalid value")
	assert.Equal(t, 64, solution(128, 2), "invalid value")
	assert.Equal(t, 0, solution(3, 5), "invalid value")
	assert.Equal(t, 1, solution(5, 5), "invalid value")
}
