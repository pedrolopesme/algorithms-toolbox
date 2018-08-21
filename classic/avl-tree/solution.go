package main

// AvlTree defines the basic tree structure
type AvlTree struct {
	root   *Node
	height int
}

// Node definition used throughout the code.
type Node struct {
	id     int
	height int
	left   *Node
	right  *Node
}

// CalcHeight returns tree's height
// TODO add tests
func (tree AvlTree) CalcHeight() int {
	if tree.root == nil {
		return 0
	} else {
		return tree.root.height
	}
}

// IsLeaf checks if a node is leafs
// TODO add tests
func (node Node) IsLeaf() bool {
	return node.height == 0
}

// max return the maximum value between two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// RotateRight is a helper that rotates node positions
// in a clockwise direction
// TODO add tests
func RotateRight(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	leftNode := node.left
	subNode := leftNode.right

	leftNode.right = node
	node.left = subNode

	node.height = max(node.left.height, node.right.height) + 1
	node.height = max(leftNode.left.height, leftNode.right.height) + 1
	return leftNode
}
