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

// GetIds is a utility func that
// visits all nodes and return theirs ids
// TODO add tests
func GetIds(node *Node) []int {
	if node != nil {
		current := []int{node.id}

		left := GetIds(node.left)
		if left != nil {
			current = append(current, left...)
		}

		right := GetIds(node.right)
		if right != nil {
			current = append(current, right...)
		}
		return current
	}
	return nil
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

// Insert adds a node to a tree
// TODO add tests
func (tree AvlTree) Insert(newNode *Node) {
	if tree.root == nil {
		tree.root = newNode
	} else {
		Append(tree.root, newNode)
	}
}

// Append appends a node to another node
// TODO add tests
func Append(node *Node, newNode *Node) *Node {
	if node == nil {
		return newNode
	}

	if node.id < newNode.id {
		node.left = Append(node.left, newNode)
	} else if node.id > newNode.id {
		node.right = Append(node.right, newNode)
	} else {
		return node
	}

	node.height = max(node.left.height, node.right.height) + 1

	// Balancing the tree
	balance := node.calcBalance()
	if balance > 1 && newNode.id < node.left.id {
		return RotateRight(node)
	}

	if balance < -1 && newNode.id > node.right.id {
		return RotateLeft(node)
	}

	if balance > 1 && newNode.id > node.left.id {
		node.left = RotateLeft(node.left)
		return RotateRight(node)
	}

	if balance < -1 && newNode.id < node.right.id {
		node.right = RotateRight(node.right)
		return RotateLeft(node)
	}

	return node
}

// IsLeaf checks if a node is leafs
// TODO add tests
func (node Node) IsLeaf() bool {
	return node.height == 0
}

// calcBalance get the balance factor of a node
// TODO add tests
func (node Node) calcBalance() int {
	if node.left == nil && node.right == nil {
		return 0
	} else if node.left == nil && node.right != nil {
		return node.right.height
	} else if node.left != nil && node.right == nil {
		return node.left.height
	} else {
		return node.left.height - node.right.height
	}
}

// max return the maximum value between two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// RotateLeft is a helper that rotates node positions
// in a counterclockwise direction
// TODO add tests
func RotateLeft(node *Node) *Node {
	if node == nil || node.right == nil {
		return node
	}

	rightNode := node.right
	rightNode.left = node
	node.right = rightNode.left

	node.height = max(node.left.height, node.right.height) + 1
	rightNode.height = max(rightNode.left.height, rightNode.right.height) + 1
	return rightNode
}

// RotateRight is a helper that rotates node positions
// in a clockwise direction
// TODO add tests
func RotateRight(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	leftNode := node.left
	leftNode.right = node
	node.left = leftNode.right

	node.height = max(node.left.height, node.right.height) + 1
	leftNode.height = max(leftNode.left.height, leftNode.right.height) + 1
	return leftNode
}
