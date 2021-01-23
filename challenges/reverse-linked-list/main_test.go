package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_WhenPrependsANodeIntoAEmptyList_NodeShouldBePrepended(t *testing.T) {
	node := &Node{val: 1}
	list := List{}
	assert.Empty(t, list.size)
	list.Prepend(node)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.val)
}

func TestList_WhenPrependsANodeIntoAList_NodeShouldBePrepended(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	list := List{}

	list.Prepend(node1)
	list.Prepend(node2)
	list.Prepend(node3)
	assert.Equal(t, 3, list.size)
	assert.Equal(t, 3, list.head.val)
	assert.Equal(t, 2, list.head.next.val)
	assert.Equal(t, 1, list.head.next.next.val)
}

func TestList_WhenoAppendsANodeIntoAEmptyList_NodeShouldBeAppended(t *testing.T) {
	node := &Node{val: 1}
	list := List{}
	assert.Empty(t, list.size)
	list.Append(node)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.val)
}

func TestList_WhenAppendsANodeIntoAList_NodeShouldBeAppended(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	list := List{}

	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	assert.Equal(t, 3, list.size)
	assert.Equal(t, 1, list.head.val)
	assert.Equal(t, 2, list.head.next.val)
	assert.Equal(t, 3, list.head.next.next.val)
}

func TestList_WhenInsertsANodeIntoAEmptyList_NodeShouldBeInserted(t *testing.T) {
	node := &Node{val: 1}
	list := List{}
	assert.Empty(t, list.size)
	list.Insert(99, node)
	assert.Equal(t, 1, list.size)
	assert.Equal(t, 1, list.head.val)
}

func TestList_WhenInsertsANodeIntoAPositionBiggerThanTheListSize_NodeShouldBeAppendedAtLastPosition(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	list := List{}
	list.Append(node1)     // <- List Right Now with only one node at first position
	list.Insert(99, node2) // <- 99 is a position bigger than the list total size
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 1, list.head.val)
	assert.Equal(t, 2, list.tail.val)
}

func TestList_WhenInsertsANodeIntoTheFirstPosition_NodeShouldBePrepended(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	list := List{}
	list.Append(node1)
	list.Insert(0, node2)
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 2, list.head.val)
	assert.Equal(t, 1, list.tail.val)
}

func TestList_WhenInsertsANodeIntoTheMiddle_NodeShouldBeInserted(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	list := List{}
	list.Prepend(node1)
	list.Append(node3)
	list.Insert(1, node2)
	assert.Equal(t, 3, list.size)
	assert.Equal(t, node1.val, list.head.val)
	assert.Equal(t, node3.val, list.tail.val)
	assert.Equal(t, node2.val, list.head.next.val) // our middle guy
}

func TestList_ReverseTheCurrentList_ReturnsANewListInverted(t *testing.T) {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	list := List{}
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	reversedList := list.Reverse()
	assert.Equal(t, 3, reversedList.size)
	assert.Equal(t, 3, reversedList.head.val)
	assert.Equal(t, 2, reversedList.head.next.val)
	assert.Equal(t, 1, reversedList.head.next.next.val)
}
