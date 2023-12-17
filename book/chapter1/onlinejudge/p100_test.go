package onlinejudge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 20, maximumCycle(1, 10), "invalid value")
	assert.Equal(t, 125, maximumCycle(100, 200), "invalid value")
	assert.Equal(t, 125, maximumCycle(100, 200), "invalid value")
	assert.Equal(t, 89, maximumCycle(201, 210), "invalid value")
	assert.Equal(t, 174, maximumCycle(900, 1000), "invalid value")
}
