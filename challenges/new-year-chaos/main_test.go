package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinimumBribes_BaseCases(t *testing.T) {
	assert.Equal(t, int32(3), minimumBribes([]int32{2, 1, 5, 3, 4}))
}
