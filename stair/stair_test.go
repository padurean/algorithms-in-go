package stair_test

import (
	"testing"

	"github.com/padurean/algorithms-in-go/stair"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, int64(1), stair.Solution(0), "should return 1 for n = 0")

	// 1 -> 1 -> 1
	// 1 -> 2
	// 2 -> 1
	// 3
	assert.Equal(t, int64(4), stair.Solution(3), "should return 4 for n = 3")

	// 1 -> 1 -> 1 -> 1
	// 1 -> 1 -> 2
	// 1 -> 2 -> 1
	// 2 -> 1 -> 1
	// 2 -> 2
	// 1 -> 3
	// 3 -> 1
	assert.Equal(t, int64(7), stair.Solution(4), "should return 7 for n = 4")

	// 1 -> 1 -> 1 -> 1 -> 1
	// 1 -> 1 -> 1 -> 2
	// 1 -> 1 -> 2 -> 1
	// 1 -> 2 -> 1 -> 1
	// 2 -> 1 -> 1 -> 1
	// 2 -> 2 -> 1
	// 2 -> 1 -> 2
	// 1 -> 2 -> 2
	// 1 -> 1 -> 3
	// 1 -> 3 -> 1
	// 3 -> 1 -> 1
	// 2 -> 3
	// 3 -> 2
	assert.Equal(t, int64(13), stair.Solution(5), "should return 13 for n = 5")

	assert.Equal(t, int64(121415), stair.Solution(20), "should return 121415 for n = 20")
	assert.Equal(t, int64(410744), stair.Solution(22), "should return 410744 for n = 22")
	assert.Equal(t, int64(501774317241), stair.Solution(45), "should return 501774317241 for n = 45")
}
