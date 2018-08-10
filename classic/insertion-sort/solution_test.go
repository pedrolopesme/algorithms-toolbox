package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSortOnUnSortedArray(test *testing.T) {
	input := []int{0,6,3,8,4,1,7}
	expectedOutput := []int{0,1,3,4,6,7,8}
	assert.Equal(test, expectedOutput, InsertionSort(input))
}

func TestInsertionSortOnSortedArray(test *testing.T) {
	input := []int{0,1,3,4,6,7,8}
	expectedOutput := []int{0,1,3,4,6,7,8}
	assert.Equal(test, expectedOutput, InsertionSort(input))
}
