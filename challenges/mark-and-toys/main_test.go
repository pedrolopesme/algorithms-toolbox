package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaximumToys_BaseCases(t *testing.T) {
	assert.Equal(t, int32(3), maximumToys([]int32{1, 2, 3, 4}, int32(7)))
}
