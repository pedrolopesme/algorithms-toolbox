package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatFlavors_Case1(t *testing.T) {
	assert.Equal(t, []int32{1, 4}, whatFlavors([]int32{1, 4, 5, 3, 2}, 4))
}

func TestWhatFlavors_Case2(t *testing.T) {
	assert.Equal(t, []int32{1, 2}, whatFlavors([]int32{2, 2, 4, 3}, 4))
}
