package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSRecursive_Case1(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, 6, binarySearchRecursive(arr, 6, 0, len(arr)-1))
}

func TestBSRecursive_Case2(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, -1, binarySearchRecursive(arr, 10, 0, len(arr)-1))
}
