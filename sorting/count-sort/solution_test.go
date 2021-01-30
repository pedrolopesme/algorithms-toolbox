package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSortOnUnSortedArray(test *testing.T) {
	input := []rune{'b', 'a', 'a', 'y', 'f', 'u'}
	expectedOutput := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	assert.Equal(test, expectedOutput, CountSort(input))
}

func TestCountSortOnSortedArray(test *testing.T) {
	input := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	expectedOutput := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	assert.Equal(test, expectedOutput, CountSort(input))
}
