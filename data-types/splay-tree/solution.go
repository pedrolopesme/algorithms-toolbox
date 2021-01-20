package main

// SplayTree defines the basic tree structure
type SplayTree struct {
	root *Node
}

// Node definition used throughout the code.
type Node struct {
	id    int
	left  *Node
	right *Node
}

var deepestLevel = 0
var deepestNode *Node

// calculateDeepestNode calculates the deepest node recursively
func calculateDeepestNode(node *Node, level int) {
	if node != nil {
		level++
		calculateDeepestNode(node.left, level)
		if level > deepestLevel {
			deepestNode = node
			deepestLevel = level
		}
		level++
		calculateDeepestNode(node.right, level)
	}
}

// FindDeepest returns the deepest node if it was found
func (tree *SplayTree) FindDeepest() *Node {
	deepestLevel = 0
	deepestNode = nil
	calculateDeepestNode(tree.root, 0)
	return deepestNode
}

// SplayDeepest finds the deepest node in the
// tree and splay it to root
func (tree *SplayTree) SplayDeepest() {
	node := tree.FindDeepest()
	if node != nil {
		tree.Get(node.id)
	}
}

// Get search for a node by its key and return it's value.
func (tree *SplayTree) Get(id int) *Node {
	node := Splay(tree.root, id)
	if node != nil {
		tree.root = node
		if tree.root.id == id {
			return tree.root
		}
	}
	return nil
}

// Add adds a node to the tree and Splay it to root node.
func (tree *SplayTree) Add(id int) {
	if tree.root == nil {
		tree.root = &Node{id: id}
		return
	}

	tree.root = Splay(tree.root, id)

	if id < tree.root.id {
		node := &Node{id: id}
		node.left = tree.root.left
		node.right = tree.root
		tree.root.left = nil
		tree.root = node

	} else if id > tree.root.id {
		node := &Node{id: id}
		node.right = tree.root.right
		node.left = tree.root
		tree.root.right = nil
		tree.root = node
	} else {
		tree.root.id = id
	}
}

// Remove removes a node from the tree
func (tree *SplayTree) Remove(id int) {
	if tree.root == nil {
		return
	}

	if tree.root.id == id {
		if tree.root.left == nil {
			tree.root = tree.root.right
		} else {
			node := tree.root.right
			tree.root = tree.root.left
			Splay(tree.root, id)
			tree.root.right = node
		}
	}
}

// GetSize is a helper function that calculates tree size.
func (tree *SplayTree) GetSize() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.GetSize()
}

// GetSize is a helper function that calculates the size
// from a specific node
func (node *Node) GetSize() int {
	leftSize := func() int {
		if node.left == nil {
			return 0
		} else {
			return node.left.GetSize()
		}
	}
	rightSize := func() int {
		if node.right == nil {
			return 0
		} else {
			return node.right.GetSize()
		}
	}
	return 1 + leftSize() + rightSize()
}

// Splay splays a node to the root of the tree. If there isn't a node with
// that key, the last node along the search path for the key will be splayed to
// the root.
func Splay(node *Node, id int) *Node {
	if node == nil {
		return nil
	}

	if id < node.id {
		// Key not found in tree
		if node.left == nil {
			return node
		}

		if id < node.left.id {
			node.left.left = Splay(node.left.left, id)
			node = RotateRight(node)
		} else if id > node.left.id {
			node.left.right = Splay(node.left.right, id)
			if node.left.right != nil {
				node.left = RotateLeft(node.left)
			}
		}

		if node.left != nil {
			node = RotateRight(node)
		}

	} else if id > node.id {
		// Key not found in tree
		if node.right == nil {
			return node
		}

		if id < node.right.id {
			node.right.left = Splay(node.right.left, id)
			if node.right.left != nil {
				node.right = RotateRight(node.right)
			}
		} else if id > node.right.id {
			node.right.right = Splay(node.right.right, id)
			node = RotateLeft(node)
		}

		if node.right != nil {
			node = RotateLeft(node)
		}
	}

	return node
}

// RotateLeft is a helper that rotates node positions in a counterclockwise direction
func RotateLeft(node *Node) *Node {
	if node == nil || node.right == nil {
		return node
	}

	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node
	return rightNode
}

// RotateRight is a helper that rotates node positions in a clockwise direction
func RotateRight(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	leftNode := node.left
	node.left = leftNode.left
	leftNode.right = node
	return leftNode
}
