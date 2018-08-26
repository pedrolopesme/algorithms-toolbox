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

	assert.Nil(test, rotatedNode.right)
	assert.Nil(test, rotatedNode.left.right)
}

func TestRotateRightWithSingleNode(test *testing.T) {
	node := &Node{id: 10}
	rotatedNode := RotateRight(node)
	assert.Equal(test, rotatedNode.id, node.id)
	assert.Nil(test, rotatedNode.left)
	assert.Nil(test, rotatedNode.right)
}

func TestRotateRightWithNodeOnTheLeftOnly(test *testing.T) {
	node := &Node{
		id:   2,
		left: &Node{id: 1, height: 1},
	}
	rotatedNode := RotateRight(node)
	assert.Equal(test, 1, rotatedNode.id)
	assert.Equal(test, 2, rotatedNode.right.id)
}

func TestRotateRightWithNodeOnTheRightOnly(test *testing.T) {
	node := &Node{
		id:    1,
		right: &Node{id: 2, height: 1},
	}
	rotatedNode := RotateRight(node)
	assert.Equal(test, 1, rotatedNode.id)
	assert.Equal(test, 2, rotatedNode.right.id)
}

func TestRotateRightWithNodesOnBothSides(test *testing.T) {
	node := &Node{
		id:    2,
		left:  &Node{id: 1, height: 1},
		right: &Node{id: 3, height: 1},
	}
	rotatedNode := RotateRight(node)

	assert.Equal(test, 1, rotatedNode.id)
	assert.Equal(test, 2, rotatedNode.right.id)
	assert.Equal(test, 3, rotatedNode.right.right.id)
	assert.Nil(test, rotatedNode.left)
	assert.Nil(test, rotatedNode.right.left)
}

func TestCalcBalanceWithNoSubNodes(test *testing.T) {
	node := &Node{
		id: 1,
	}
	balance := node.calcBalance()
	assert.Equal(test, 0, balance)
}

func TestCalcBalanceWithNodesOnLeftSide(test *testing.T) {
	node := &Node{
		id:   1,
		left: &Node{id: 2, height: 1},
	}
	balance := node.calcBalance()
	assert.Equal(test, 1, balance)
}

func TestCalcBalanceWithNodesOnRightSide(test *testing.T) {
	node := &Node{
		id:    1,
		right: &Node{id: 2, height: 1},
	}
	balance := node.calcBalance()
	assert.Equal(test, 1, balance)
}

func TestCalcBalanceWithNodesOnBothSidesAndTheSameHeight(test *testing.T) {
	node := &Node{
		id:    1,
		left:  &Node{id: 2, height: 1},
		right: &Node{id: 2, height: 1},
	}
	balance := node.calcBalance()
	assert.Equal(test, 0, balance)
}

func TestCalcBalanceWithNodesOnBothSidesAndDifferentHeights(test *testing.T) {
	node := &Node{
		id:    1,
		left:  &Node{id: 2, height: 2},
		right: &Node{id: 2, height: 1},
	}
	balance := node.calcBalance()
	assert.Equal(test, 1, balance)
}

func TestCalcHeightOnEmptyTree(test *testing.T) {
	tree := &AvlTree{}
	assert.Equal(test, 0, tree.CalcHeight())
}

func TestCalcHeightOnNonEmptyTree(test *testing.T) {
	tree := &AvlTree{root: &Node{height: 1}}
	assert.Equal(test, 1, tree.CalcHeight())
}

func TestGetIdsOnEmptyTree(test *testing.T) {
	var node *Node
	var expected []int
	assert.Equal(test, expected, GetIds(node))
}

func TestGetIdsWithSingleNode(test *testing.T) {
	node := &Node{id: 1}
	expected := []int{1}
	assert.Equal(test, expected, GetIds(node))
}

func TestGetIdsWithNodesOnTheLeft(test *testing.T) {
	node := &Node{
		id:   1,
		left: &Node{id: 2, left: &Node{id: 3}},
	}
	expected := []int{1, 2, 3}
	assert.Equal(test, expected, GetIds(node))
}

func TestGetIdsWithNodesOnTheRight(test *testing.T) {
	node := &Node{
		id:    1,
		right: &Node{id: 2, right: &Node{id: 3}},
	}
	expected := []int{1, 2, 3}
	assert.Equal(test, expected, GetIds(node))
}

func TestGetIdsWithNodesOnBothSide(test *testing.T) {
	node := &Node{
		id:    3,
		left:  &Node{id: 2, left: &Node{id: 1}},
		right: &Node{id: 4, right: &Node{id: 5}},
	}
	expected := []int{3, 2, 1, 4, 5}
	assert.Equal(test, expected, GetIds(node))
}

func TestAppendOnNodeWithNodeSubNodes(test *testing.T) {
	node := &Node{id: 1}
	newNode := &Node{id: 2}
	appendNode := Append(node, newNode)

	expected := []int{1, 2}
	assert.Equal(test, expected, GetIds(appendNode))
}

func TestAppendToTheRightOnNodeWithNodeLeftSubNodes(test *testing.T) {
	node := &Node{id: 3, left: &Node{id: 1}}
	newNode := &Node{id: 5}
	appendNode := Append(node, newNode)

	expected := []int{3, 1, 5}
	assert.Equal(test, expected, GetIds(appendNode))
}

func TestAppendToTheLeftOnNodeWithNodeLeftSubNodes(test *testing.T) {
	node := &Node{id: 3, left: &Node{id: 2}}
	newNode := &Node{id: 1}
	appendNode := Append(node, newNode)

	expected := []int{3, 2, 1}
	assert.Equal(test, expected, GetIds(appendNode))
}
func TestAppendToTheRightOnNodeWithNodeRightSubNodes(test *testing.T) {
	node := &Node{id: 3, right: &Node{id: 5}}
	newNode := &Node{id: 7}
	appendNode := Append(node, newNode)

	expected := []int{3, 5, 7}
	assert.Equal(test, expected, GetIds(appendNode))
}

func TestAppendToTheLeftOnNodeWithNodeRightSubNodes(test *testing.T) {
	node := &Node{id: 3, right: &Node{id: 5}}
	newNode := &Node{id: 4}
	appendNode := Append(node, newNode)

	expected := []int{3, 4, 5}
	assert.Equal(test, expected, GetIds(appendNode))
}

func TestAppendToTheLeftOnNodeWithNodesOnBothSides(test *testing.T) {
	node := &Node{
		id:    3,
		left:  &Node{id: 2},
		right: &Node{id: 4},
	}
	newNode := &Node{id: 1}
	appendNode := Append(node, newNode)

	expected := []int{3, 2, 4, 1}
	assert.Equal(test, expected, GetIds(appendNode))
}

func TestAppendToTheRightOnNodeWithNodesOnBothSides(test *testing.T) {
	node := &Node{
		id:    3,
		left:  &Node{id: 2},
		right: &Node{id: 4},
	}
	newNode := &Node{id: 5}
	appendNode := Append(node, newNode)

	expected := []int{3, 2, 5, 4}
	assert.Equal(test, expected, GetIds(appendNode))
}
