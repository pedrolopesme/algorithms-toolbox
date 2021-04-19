package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAbsDiff_Case1(t *testing.T) {
	assert.Equal(t, int32(3), minimumAbsoluteDifference([]int32{3, -7, 0}))
}

func TestMinAbsDiff_Case2(t *testing.T) {
	assert.Equal(t, int32(1), minimumAbsoluteDifference([]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}))
}

func TestMinAbsDiff_Case3(t *testing.T) {
	assert.Equal(t, int32(3), minimumAbsoluteDifference([]int32{1, -3, 71, 68, 17}))
}
