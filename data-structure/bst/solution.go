package main

// Tree defines the basic tree structure
type Tree struct {
	root *Node
}

// BST specifies a Binary Search Tree
type BST interface {
	Search(id int) *Node
	Insert(id int)
	Delete(id int)
}

// Node definition used throughout the code.
type Node struct {
	id    int
	left  *Node
	right *Node
}

// NewNode is a node builder
func NewNode(id int) *Node {
	return &Node{id: id}
}

// Search tries to find a node by its key
func (tree Tree) Search(id int) *Node {
	return Search(tree.root, id)
}

// Insert adds a node direct to a tree
func (tree *Tree) Insert(id int) {
	if tree.root == nil {
		tree.root = &Node{id: id}
	} else {
		tree.root = Insert(tree.root, id)
	}
}

// Delete removes a node from the tree
func (tree *Tree) Delete(id int) {
	tree.root = Delete(tree.root, id)
}

// Insert adds a node to a tree
func Insert(root *Node, id int) *Node {
	if root == nil {
		root = NewNode(id)
		return root
	}

	if id < root.id {
		root.left = Insert(root.left, id)
	} else if id > root.id {
		root.right = Insert(root.right, id)
	}

	return root
}

func Delete(root *Node, id int) *Node {
	if root == nil {
		return root
	}

	if id < root.id {
		root.left = Delete(root.left, id)
	} else if id > root.id {
		root.right = Delete(root.right, id)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		root.id = minValue(root.right)
		root.right = Delete(root.right, root.id)
	}

	return root
}

// minValueNode returns the minimum value of a tree
func minValue(root *Node) int {
	minValue := root.id
	for root.left != nil {
		minValue = root.left.id
		root = root.left
	}
	return minValue
}

// Insert adds a node to a tree
func Search(root *Node, id int) *Node {
	if root == nil || root.id == id {
		return root
	}

	if id < root.id {
		return Search(root.left, id)
	}

	return Search(root.right, id)
}