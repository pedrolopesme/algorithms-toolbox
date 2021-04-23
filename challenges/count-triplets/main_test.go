package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountTriplet_Case1(t *testing.T) {
	arr := []int64{1, 4, 16, 64}
	assert.Equal(t, int64(2), countTriplets(arr, int64(4)))
}

func TestCountTriplet_Case2(t *testing.T) {
	arr := []int64{1, 2, 2, 4}
	assert.Equal(t, int64(2), countTriplets(arr, int64(2)))
}

func TestCountTriplet_Case3(t *testing.T) {
	arr := []int64{1, 3, 9, 9, 27, 81}
	assert.Equal(t, int64(6), countTriplets(arr, int64(3)))
}

func TestCountTriplet_Case4(t *testing.T) {
	arr := []int64{1, 5, 5, 25, 125}
	assert.Equal(t, int64(4), countTriplets(arr, int64(5)))
}

func TestCountTriplet_Case5(t *testing.T) {
	arr := []int64{1, 1, 1, 1}
	assert.Equal(t, int64(3), countTriplets(arr, int64(1)))
}
