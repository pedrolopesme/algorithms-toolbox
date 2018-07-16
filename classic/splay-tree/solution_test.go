package main

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestGetANonExistentNode(test *testing.T) {
	tree := SplayTree{}
	tree.Add(100)
	assert.Nil(test, tree.Get(1))
}

func TestAddAndGetANode(test *testing.T) {
	tree := SplayTree{}
	tree.Add(100)
	assert.Equal(test, 100, tree.Get(100).id)
}

func TestAddAndGetMultipleNodes(test *testing.T) {
	tree := SplayTree{}
	tree.Add(100)
	tree.Add(101)
	tree.Add(102)
	assert.Equal(test, 100, tree.Get(100).id)
	assert.Equal(test, 101, tree.Get(101).id)
	assert.Equal(test, 102, tree.Get(102).id)
}

func TestTreeSizeWithNoNodes(test *testing.T) {
	tree := SplayTree{}
	assert.Equal(test, 0, tree.GetSize())
}

func TestTreeSizeWithOneNode(test *testing.T) {
	tree := SplayTree{}
	tree.Add(100)
	assert.Equal(test, 1, tree.GetSize())
}

func TestTreeSizeWithMultipleNodes(test *testing.T) {
	tree := SplayTree{}
	tree.Add(101)
	tree.Add(102)
	tree.Add(103)
	assert.Equal(test, 3, tree.GetSize())
}


func TestRotateLeftAnEmptyNode(test *testing.T) {
	assert.Nil(test, RotateLeft(nil))
}

func TestRotateLeftANodeWithNoLeafs(test *testing.T) {
	node := Node{id: 100}
	assert.Equal(test, node.id, RotateLeft(&node).id)
}
