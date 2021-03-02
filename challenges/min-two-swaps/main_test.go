package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinSwaps_BaseCases(t *testing.T) {
	assert.Equal(t, int32(5), minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}))
	assert.Equal(t, int32(3), minimumSwaps([]int32{4, 3, 1, 2}))
}
