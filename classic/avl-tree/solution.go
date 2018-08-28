package main

// AvlTree defines the basic tree structure
type AvlTree struct {
	root *Node
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
func GetIds(node *Node) (current []int) {
	if node != nil {
		current = []int{node.id}

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
	return
}

// CalcHeight returns tree's height
func (tree AvlTree) CalcHeight() int {
	if tree.root == nil {
		return 0
	} else {
		return tree.root.height
	}
}

// Insert adds a node to a tree
func (tree *AvlTree) Insert(newNode *Node) {
	if tree.root == nil {
		tree.root = newNode
	} else {
		tree.root = Append(tree.root, newNode)
	}
}

// Append appends a node to another node
func Append(node *Node, newNode *Node) *Node {
	if node == nil {
		return newNode
	}

	if node.id > newNode.id {
		node.left = Append(node.left, newNode)
	} else if node.id < newNode.id {
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
func (node *Node) calcBalance() int {
	leftHeight := 0
	rightHeight := 0

	if node.right != nil {
		rightHeight = node.right.height
	}

	if node.left != nil {
		leftHeight = node.left.height
	}

	return leftHeight - rightHeight
}

// MaxHeight return the maximum height value between two nodes
func MaxHeight(n1, n2 *Node) int {
	n1Height := 1
	n2Height := 1

	if n1 != nil {
		n1Height = n1.height
	}

	if n2 != nil {
		n2Height = n2.height
	}

	if n1Height > n2Height {
		return n1Height
	}
	return n2Height
}

// RotateLeft is a helper that rotates node positions
// in a counterclockwise direction
func RotateLeft(node *Node) *Node {
	if node == nil || node.right == nil {
		return node
	}

	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node

	node.height = MaxHeight(node.left, node.right) + 1
	rightNode.height = MaxHeight(rightNode.left, rightNode.right) + 1
	return rightNode
}

// RotateRight is a helper that rotates node positions
// in a clockwise direction
func RotateRight(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	leftNode := node.left
	node.left = leftNode.right
	leftNode.right = node

	node.height = MaxHeight(node.left, node.right) + 1
	leftNode.height = MaxHeight(leftNode.left, leftNode.right) + 1
	return leftNode
}

// minValueNode returns the node with minimum id
func minValueNode(node *Node) *Node {
	current := node

	/* loop down to find the leftmost leaf */
	for current.left != nil {
		current = current.left
	}

	return current
}

// DeleteNode removes a node from a tree by its ID
func DeleteNode(root *Node, id int) *Node {

	if root == nil {
		return root
	}

	if id < root.id {
		root.left = DeleteNode(root.left, id)
	} else if id > root.id {
		root.right = DeleteNode(root.right, id)
	} else {

		// node with only one child or no child
		if (root.left == nil) || (root.right == nil) {
			var temp *Node

			if temp == root.left {
				temp = root.right
			} else {
				temp = root.left
			}

			// No child case
			if temp == nil {
				temp = root
				root = nil
			} else {
				root = temp
			}
		} else {
			temp := minValueNode(root.right)

			// Copy the inorder successor's data to this node
			root.id = temp.id

			// Delete the inorder successor
			root.right = DeleteNode(root.right, temp.id)
		}
	}

	// If the tree had only one node then return
	if root == nil {
		return root
	}

	root.height = MaxHeight(root.left, root.right) + 1
	balance := root.calcBalance()

	// If this node becomes unbalanced, then there are 4 cases
	// Left Left Case
	if balance > 1 && root.left.calcBalance() >= 0 {
		return RotateRight(root)
	}

	// Left Right Case
	if balance > 1 && root.left.calcBalance() < 0 {
		root.left = RotateLeft(root.left)
		return RotateRight(root)
	}

	// Right Right Case
	if balance < -1 && root.right.calcBalance() <= 0 {
		return RotateLeft(root)
	}

	// Right Left Case
	if balance < -1 && root.right.calcBalance() > 0 {
		root.right = RotateRight(root.right)
		return RotateLeft(root)
	}

	return root
}
