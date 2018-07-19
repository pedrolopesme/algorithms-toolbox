package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushNodeToAEmptyList(test *testing.T) {
	list := LinkedList{}
	assert.Equal(test, uint64(1), list.size)
}

func TestIsEmptyWhenListIsEmpty(test *testing.T) {
	list := LinkedList{}
	assert.True(test, list.IsEmpty())
}

func TestIsEmptyWhenListIsNotEmpty(test *testing.T) {
	list := LinkedList{ first: &Node{} }
	assert.False(test, list.IsEmpty())
}
