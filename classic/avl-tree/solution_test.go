package main

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestMaxHeightWithTwoNilNodes(test *testing.T) {
	var n1 *Node
	var n2 *Node
	assert.Equal(test, 0, MaxHeight(n1, n2))
}

func TestMaxHeightWithOneNilNodes(test *testing.T) {
	n1 := &Node{height: 6}
	var n2 *Node
	assert.Equal(test, 6, MaxHeight(n1, n2))
}

func TestMaxHeightWithTwoPositiveNodes(test *testing.T) {
	n1 := &Node{height: 6}
	n2 := &Node{height: 60}
	assert.Equal(test, 60, MaxHeight(n1, n2))
}

func TestMaxHeightWithTwoNegativeNodes(test *testing.T) {
	n1 := &Node{height: -6}
	n2 := &Node{height: -60}
	assert.Equal(test, -6, MaxHeight(n1, n2))
}

func TestMaxHeightWithTwoEqualNodes(test *testing.T) {
	n1 := &Node{height: 60}
	n2 := &Node{height: 60}
	assert.Equal(test, 60, MaxHeight(n1, n2))
}

func TestIsLeafWithLeafNode(test *testing.T) {
	node := Node{height: 0}
	assert.True(test, node.IsLeaf())
}

func TestIsLeafWithNonLeafNode(test *testing.T) {
	node := Node{height: 10}
	assert.False(test, node.IsLeaf())
}

func TestRotateLeftWithSingleNode(test *testing.T) {
	node := &Node{id: 10}
	rotatedNode := RotateLeft(node)
	assert.Equal(test, rotatedNode.id, node.id)
	assert.Nil(test, rotatedNode.left)
	assert.Nil(test, rotatedNode.right)
}

func TestRotateLeftWithNodeOnTheLeftOnly(test *testing.T) {
	node := &Node{
		id:   2,
		left: &Node{id: 1, height: 1},
	}
	rotatedNode := RotateLeft(node)
	assert.Equal(test, 2, rotatedNode.id)
	assert.Equal(test, 1, rotatedNode.left.id)
}

func TestRotateLeftWithNodeOnTheRightOnly(test *testing.T) {
	node := &Node{
		id:    1,
		right: &Node{id: 2, height: 1},
	}
	rotatedNode := RotateLeft(node)
	assert.Equal(test, 2, rotatedNode.id)
	assert.Equal(test, 1, rotatedNode.left.id)
}

func TestRotateLeftWithNodesOnBothSides(test *testing.T) {
	node := &Node{
		id:    2,
		left:  &Node{id: 1, height: 1},
		right: &Node{id: 3, height: 1},
	}
	rotatedNode := RotateLeft(node)
	assert.Equal(test, 3, rotatedNode.id)
	assert.Equal(test, 2, rotatedNode.left.id)
	assert.Equal(test, 1, rotatedNode.left.left.id)
}
