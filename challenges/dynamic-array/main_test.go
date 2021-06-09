package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray_BaseCase(t *testing.T) {
	assert.Equal(t,
		[]int32{7, 3},
		dynamicArray(int32(2), [][]int32{
			[]int32{1, 0, 5},
			[]int32{1, 1, 7},
			[]int32{1, 0, 3},
			[]int32{2, 1, 0},
			[]int32{2, 1, 1},
		}),
	)
}
