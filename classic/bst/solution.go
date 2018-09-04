package main

// Tree defines the basic tree structure
type Tree struct {
	root *Node
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
// TODO implement and add tests
func (t Tree) Search(id int) *Node {
	return nil
}

// GetIds is a utility func that
// visits all nodes and return their ids
// TODO add tests
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

// Insert adds a node direct to a tree
func (tree *Tree) Insert(id int) {
	if tree.root == nil {
		tree.root = &Node{id: id}
	} else {
		tree.root = Insert(tree.root, id)
	}
}

// Insert adds a node to a tree
// TODO add tests
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
