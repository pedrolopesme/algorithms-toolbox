package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestHourGlass_BaseCaseWithPositiveNumbers(t *testing.T) {
	arr := make([][]int32, 6)
	arr[0] = append(arr[0], 1, 1, 1, 0, 0, 0)
	arr[1] = append(arr[1], 0, 1, 0, 0, 0, 0)
	arr[2] = append(arr[2], 1, 1, 1, 0, 0, 0)
	arr[3] = append(arr[3], 0, 0, 2, 4, 4, 0)
	arr[4] = append(arr[4], 0, 0, 0, 2, 0, 0)
	arr[5] = append(arr[5], 0, 0, 1, 2, 4, 0)
	assert.Equal(t, int32(19), hourglassSum(arr))
}

func TestHourGlass_BaseCaseWithNegativeNumbers(t *testing.T) {
	arr := make([][]int32, 6)
	arr[0] = append(arr[0], -1, -1, 0, -9, -2, -2)
	arr[1] = append(arr[1], -2, -1, -6, -8, -2, -5)
	arr[2] = append(arr[2], -1, -1, -1, -2, -3, -4)
	arr[3] = append(arr[3], -1, -9, -2, -4, -4, -5)
	arr[4] = append(arr[4], -7, -3, -3, -2, -9, -9)
	arr[5] = append(arr[5], -1, -3, -1, -2, -4, -5)
	assert.Equal(t, int32(-6), hourglassSum(arr))
}
