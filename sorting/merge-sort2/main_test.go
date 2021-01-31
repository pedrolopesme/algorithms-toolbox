package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort_BaseCases(t *testing.T) {
	given := []int{5, 6, 4, 3, 8, 9, 2, 1, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, mergeSort(given))
}
