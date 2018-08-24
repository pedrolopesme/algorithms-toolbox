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
// visits all nodes and return their ids
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

	node.height = MaxHeight(node.left, node.right) + 1

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

// MaxHeight return the maximum height value between two nodes
func MaxHeight(n1, n2 *Node) (maxVal int) {
	if n1 == nil && n2 == nil {
		maxVal = 0
	} else if n1 == nil && n2 != nil {
		maxVal = n2.height
	} else if n1 != nil && n2 == nil {
		maxVal = n1.height
	} else if n1.height > n2.height {
		maxVal = n1.height
	} else if n1.height < n2.height {
		maxVal = n2.height
	} else {
		maxVal = n2.height
	}
	return maxVal
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

	node.height = MaxHeight(node.left, node.right) + 1
	rightNode.height = MaxHeight(rightNode.left, rightNode.right) + 1
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

	node.height = MaxHeight(node.left, node.right) + 1
	leftNode.height = MaxHeight(leftNode.left, leftNode.right) + 1
	return leftNode
}
