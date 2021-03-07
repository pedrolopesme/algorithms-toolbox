package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinimumBribes_BaseCases(t *testing.T) {
	assert.Equal(t, int32(1), minimumBribes([]int32{1, 2, 3, 5, 4, 6, 7, 8}))
	assert.Equal(t, int32(-1), minimumBribes([]int32{4, 1, 2, 3}))
	assert.Equal(t, int32(3), minimumBribes([]int32{2, 1, 5, 3, 4}))
	assert.Equal(t, int32(7), minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4}))
}
