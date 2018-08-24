package main

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestMaxWithTwoNegativeNumbers(test *testing.T) {
	assert.Equal(test, -6, max(-6, -60))
}

func TestMaxWithTwoPositiveNumbers(test *testing.T) {
	assert.Equal(test, 60, max(6, 60))
}

func TestMaxWithTwoEqualNumbers(test *testing.T) {
	assert.Equal(test, 60, max(60, 60))
}

func TestIsLeafWithLeafNode(test *testing.T) {
	node := Node{height: 0}
	assert.True(test, node.IsLeaf())
}

func TestIsLeafWithNonLeafNode(test *testing.T) {
	node := Node{height: 10}
	assert.False(test, node.IsLeaf())
}
