package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayManipulation_BaseCases(t *testing.T) {
	assert.Equal(t, int64(10), arrayManipulation(int32(10), [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}))
	assert.Equal(t, int64(200), arrayManipulation(int32(5), [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
	assert.Equal(t, int64(882), arrayManipulation(int32(4), [][]int32{{2, 3, 603}, {1, 1, 286}, {4, 4, 882}}))
}

func BenchmarkArrayManipulation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrayManipulation(int32(10), [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}})
	}
}

func BenchmarkArrayManipulation2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrayManipulation2(int32(10), [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}})
	}
}
