package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSortOnUnSortedArray(test *testing.T) {
	input := []rune{'b', 'a', 'a', 'y', 'f', 'u'}
	expectedOutput := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	assert.Equal(test, expectedOutput, CountSortRunes(input))
}

func TestCountSortOnSortedArray(test *testing.T) {
	input := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	expectedOutput := []rune{'a', 'a', 'b', 'f', 'u', 'y'}
	assert.Equal(test, expectedOutput, CountSortRunes(input))
}

func TestCountSortInt_Case1(test *testing.T) {
	input := []int{1, 4, 1, 2, 7, 5, 2}
	expectedOutput := []int{1, 1, 2, 2, 4, 5, 7}
	assert.Equal(test, expectedOutput, CountSortInt(input))
}
