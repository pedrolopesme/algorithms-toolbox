package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArrays_WhenTwoValidArraysWereGiven_ThenItReturnsTheMergedArray(t *testing.T) {
	arr1 := []int{0, 1, 2, 3}
	arr2 := []int{4, 5, 6}
	expected := []int{0, 1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, MergeSortedArrays(arr1, arr2))
}
