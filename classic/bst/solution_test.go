package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinValueWithNodesOnTheLeft(test *testing.T) {
	n1 := NewNode(10)
	n1.left = NewNode(9)
	n1.left.left = NewNode(8)
	assert.Equal(test, 8, minValue(n1))
}

func TestMinValueWithNodesOnTheRight(test *testing.T) {
	n1 := NewNode(5)
	n1.right = NewNode(6)
	n1.right.right = NewNode(7)
	assert.Equal(test, 5, minValue(n1))
}

func TestMinValueWithNodesOnBothSides(test *testing.T) {
	n1 := NewNode(10)
	n1.left = NewNode(9)
	n1.right = NewNode(11)
	n1.left.left = NewNode(8)
	n1.right.left = NewNode(12)
	assert.Equal(test, 8, minValue(n1))
}

func TestNodeInsertAddingNodesToTheLeft(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(3)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 4, tree.root.left.id)
	assert.Equal(test, 3, tree.root.left.left.id)
}

func TestNodeInsertAddingNodesToTheRight(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 6, tree.root.right.id)
	assert.Equal(test, 7, tree.root.right.right.id)
}

func TestNodeInsertAddingNodesToBothSides(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(3)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 6, tree.root.right.id)
	assert.Equal(test, 7, tree.root.right.right.id)
	assert.Equal(test, 4, tree.root.left.id)
	assert.Equal(test, 3, tree.root.left.left.id)
}

func TestSearchWhenNodeExistsOnTheLeft(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(3)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 6, tree.root.right.id)
	assert.Equal(test, 7, tree.root.right.right.id)
	assert.Equal(test, 4, tree.root.left.id)
	assert.Equal(test, 3, tree.root.left.left.id)

	node := tree.Search(3)
	assert.Equal(test, 3, node.id)
}

func TestSearchWhenNodeExistsOnTheRight(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(3)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 6, tree.root.right.id)
	assert.Equal(test, 7, tree.root.right.right.id)
	assert.Equal(test, 4, tree.root.left.id)
	assert.Equal(test, 3, tree.root.left.left.id)

	node := tree.Search(7)
	assert.Equal(test, 7, node.id)
}

func TestSearchWhenNodeDoesntExist(test *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(3)
	assert.Equal(test, 5, tree.root.id)
	assert.Equal(test, 6, tree.root.right.id)
	assert.Equal(test, 7, tree.root.right.right.id)
	assert.Equal(test, 4, tree.root.left.id)
	assert.Equal(test, 3, tree.root.left.left.id)

	node := tree.Search(10)
	assert.Nil(test, node)
}