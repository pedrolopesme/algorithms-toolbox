package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoArraysWithItemsInCommon_shouldReturnTrue(t *testing.T) {
	a1 := []rune{'a', 'b', 'c', 'd'}
	a2 := []rune{'c', 'b', 'a'}
	assert.True(t, containsCommonItens(a1, a2))
}

func TestTwoArraysWithoutItemsInCommon_shouldReturnFalse(t *testing.T) {
	a1 := []rune{'a', 'b'}
	a2 := []rune{'b', 'y'}
	assert.False(t, containsCommonItens(a1, a2))
}
