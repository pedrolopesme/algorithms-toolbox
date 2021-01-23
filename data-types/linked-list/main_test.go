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
