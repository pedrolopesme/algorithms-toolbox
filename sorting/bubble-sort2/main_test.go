package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestBubbleSort_WhenReceivesAnArray_ItShouldSortIt(t *testing.T) {
	given := []int{1, 5, 4, 2, 3, 6, 9, 7, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, bubbleSort(given))
}
