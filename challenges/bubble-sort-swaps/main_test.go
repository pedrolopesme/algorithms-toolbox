package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleCounter_BaseCase(t *testing.T) {
	swapsExpected := int32(3)
	firstExpected := int32(1)
	lastExpected := int32(6)
	swapsGiven, firstGiven, lastGiven := countSwaps([]int32{6, 4, 1})
	assert.Equal(t, swapsExpected, swapsGiven)
	assert.Equal(t, firstExpected, firstGiven)
	assert.Equal(t, lastExpected, lastGiven)
}
