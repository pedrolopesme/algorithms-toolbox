package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstRecurrying_WhenNoRepetionWasFound_ReturnNegativeOne(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	assert.Equal(t, -1, FindFirstRecurrying(arr))
}

func TestFindFirstRecurrying_WhenThereIsARepetion_ReturnTheRepeatedNumber(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 5}
	assert.Equal(t, 2, FindFirstRecurrying(arr))
}
