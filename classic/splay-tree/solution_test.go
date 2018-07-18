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

func TestRotateRightAnEmptyNode(test *testing.T) {
	assert.Nil(test, RotateRight(nil))
}

func TestRotateRightANodeWithNoLeafs(test *testing.T) {
	node := Node{id: 100}
	rotatedNode := RotateRight(&node)
	assert.Equal(test, node.id, rotatedNode.id)
	assert.Nil(test, rotatedNode.left)
	assert.Nil(test, rotatedNode.right)
}

func TestRotateRightAWhereThereIsNoNodesInLeftSide(test *testing.T) {
	tree := SplayTree{}
	tree.Add(2)
	tree.Add(1)
	rotatedNode := RotateRight(tree.root)
	assert.Equal(test, rotatedNode.id, tree.root.id)
	assert.Equal(test, 2, rotatedNode.right.id)
	assert.Nil(test, rotatedNode.left)
}

func TestRotateRight(test *testing.T) {
	tree := SplayTree{}
	tree.Add(3)
	tree.Add(1)
	tree.Add(2)
	rotatedNode := RotateRight(tree.root)
	assert.Equal(test, 1, rotatedNode.id)
	assert.Nil(test, rotatedNode.left)
	assert.Equal(test, 2, rotatedNode.right.id)
	assert.Equal(test, 3, rotatedNode.right.right.id)
}

func TestSplayOnAnEmptyTree(test *testing.T) {
	node := Splay(nil, 0)
	assert.Nil(test, node)
}

func TestSplayOnATreeWithOneNode(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	node := Splay(tree.root, 1)
	assert.Equal(test, 1, node.id)
}

func TestSplayOnATreeWithOnlyLeftSideNodes(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	node := Splay(tree.root, 1)
	assert.Equal(test, 1, node.id)
	assert.Nil(test,  node.left)
	assert.Equal(test, 2, node.right.id)
	assert.Equal(test, 3, node.right.right.id)
}

func TestSplayOnATreeWithOnlyRightSideNodes(test *testing.T) {
	tree := SplayTree{}
	tree.Add(3)
	tree.Add(2)
	tree.Add(1)
	node := Splay(tree.root, 3)
	assert.Equal(test, 3, node.id)
	assert.Nil(test,  node.right)
	assert.Equal(test, 2, node.left.id)
	assert.Equal(test, 1, node.left.left.id)
}

func TestSplayOnAVeryUnbalancedTree(test *testing.T) {
	tree := SplayTree{}
	for i := 1; i <= 1000; i++ {
		tree.Add(i)
	}

	node := Splay(tree.root, 1000)
	assert.Equal(test, 1000, node.id)
	assert.Equal(test, 1000, node.GetSize())
}

func TestSplayTheDeepestNodeOnAnEmptyTree(test *testing.T) {
	tree := SplayTree{}
	tree.SplayDeepest()
	assert.Nil(test, tree.root)
}

func TestSplayTheDeepestNodeToTheRoot(test *testing.T) {
	tree := SplayTree{}
	tree.Add(6)
	tree.Add(5)
	tree.Add(4)
	tree.Add(2)
	tree.Add(1)
	tree.Add(3)

	assert.Equal(test, 3, tree.root.id)
	tree.SplayDeepest()
	assert.Equal(test, 6, tree.root.id)
}

func TestGetOnAnEmptyTree(test *testing.T) {
	tree := SplayTree{}
	assert.Nil(test, tree.Get(0))
}

func TestGetOnATreeWithOneNodeWithMatch(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	assert.Equal(test, 1, tree.Get(1).id)
}

func TestGetOnATreeWithOneNodeWithoutMatch(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	assert.Nil(test, tree.Get(2))
}

func TestGetShouldSplayTheNode(test *testing.T) {
	tree := SplayTree{}
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	assert.Equal(test, 3, tree.root.id)
	tree.Get(1)
	assert.Equal(test, 1, tree.root.id)
}