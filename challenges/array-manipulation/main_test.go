package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayManipulation_BaseCases(t *testing.T) {
	assert.Equal(t, int64(10), arrayManipulation(int32(10), [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}))
	assert.Equal(t, int64(200), arrayManipulation(int32(5), [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
}
