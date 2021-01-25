package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush_WhenStackIsEmpty_ThenMyNewNodeWillBeTheHeadAndTail(t *testing.T) {
	node := Node{val: 1}
	stack := Stack{}
	assert.Empty(t, stack.size)

	stack.Push(&node)
	assert.Equal(t, 1, stack.size)
	assert.Equal(t, node.val, stack.head.val)
	assert.Equal(t, node.val, stack.tail.val)
}

func TestPush_WhenStackIsNotEmpty_ThenMyNewNodeWillBeTheTail(t *testing.T) {
	node1 := Node{val: 1}
	node2 := Node{val: 2}
	node3 := Node{val: 3}
	stack := Stack{}
	assert.Empty(t, stack.size)

	stack.Push(&node1)
	stack.Push(&node2)
	stack.Push(&node3)
	assert.Equal(t, 3, stack.size)
	assert.Equal(t, node1.val, stack.head.val)
	assert.Equal(t, node3.val, stack.tail.val)
}

func TestPop_WhenStackHasOneNode_ThenMyStackShouldBeEmpty(t *testing.T) {
	node := Node{val: 1}
	stack := Stack{}

	stack.Push(&node)
	assert.Equal(t, 1, stack.size)

	poppedStack := stack.Pop()
	assert.Equal(t, 0, stack.size)
	assert.Equal(t, node.val, poppedStack.val)
}

func TestPop_WhenStackIsNotEmpty_ThenMyStackShouldBeProperlyUpdated(t *testing.T) {
	node1 := Node{val: 1}
	node2 := Node{val: 2}
	node3 := Node{val: 3}
	stack := Stack{}

	stack.Push(&node1)
	stack.Push(&node2)
	stack.Push(&node3)
	assert.Equal(t, 3, stack.size)

	poppedNode := stack.Pop()
	assert.Equal(t, 2, stack.size)
	assert.Equal(t, node3.val, poppedNode.val)
}
