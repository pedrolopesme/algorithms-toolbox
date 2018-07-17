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
	for i := 0; i < 100; i++ {
		tree.Add(i)
	}
	assert.Equal(test, 100, tree.GetSize())
}

func TestRotateLeftAnEmptyNode(test *testing.T) {
	assert.Nil(test, RotateLeft(nil))
}

func TestRotateLeftANodeWithNoLeafs(test *testing.T) {
	node := Node{id: 100}
	rotatedNode := RotateLeft(&node)
	assert.Equal(test, node.id, rotatedNode.id)
	assert.Nil(test, rotatedNode.left)
	assert.Nil(test, rotatedNode.right)
}

func TestRotateLeftAWhereThereIsNoNodesInRightSide(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	tree.Add(2)
	rotatedNode := RotateLeft(tree.root)
	assert.Equal(test, rotatedNode.id, tree.root.id)
	assert.Equal(test, 1, rotatedNode.left.id)
	assert.Nil(test, rotatedNode.right)
}

func TestRotateLeft(test *testing.T) {
	tree := SplayTree{}
	tree.Add(3)
	tree.Add(1)
	tree.Add(2)
	rotatedNode := RotateLeft(tree.root)
	assert.Equal(test, 3, rotatedNode.id)
	assert.Nil(test, rotatedNode.right)
	assert.Equal(test, 2, rotatedNode.left.id)
	assert.Equal(test, 1, rotatedNode.left.left.id)
}
