package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRotLeft_BaseCase(t *testing.T) {
	givenArr := []int32{1, 2, 3, 4, 5}
	expectedArr := []int32{5, 1, 2, 3, 4}
	assert.Equal(t, expectedArr, rotLeft(givenArr, 4))
}
