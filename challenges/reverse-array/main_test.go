package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseArray_BaseCase(t *testing.T) {
	arr := []int32{0, 1, 2, 3, 4}
	assert.Equal(t, []int32{4, 3, 2, 1, 0}, reverseArray(arr))

}
