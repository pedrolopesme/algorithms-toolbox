package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushNodeToAnEmptyList(test *testing.T) {
	list := LinkedList{}
	list.Push(1)
	assert.False(test, list.IsEmpty())
	assert.Equal(test, 1, list.first.id)
}

func TestPushNodeToANonEmptyList(test *testing.T) {
	list := LinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	assert.False(test, list.IsEmpty())
	assert.Equal(test, 3, list.first.id)
	assert.Equal(test, 2, list.first.next.id)
	assert.Equal(test, 1, list.first.next.next.id)
}

func TestIsEmptyWhenListIsEmpty(test *testing.T) {
	list := LinkedList{}
	assert.True(test, list.IsEmpty())
}

func TestIsEmptyWhenListIsNotEmpty(test *testing.T) {
	list := LinkedList{ first: &Node{} }
	assert.False(test, list.IsEmpty())
}
