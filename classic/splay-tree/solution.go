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

// calculateDeepestNode calculates the deepest node recursively
func calculateDeepestNode(node *Node, level int, deepestLevel *int, deepestNode *Node) {
	if node != nil {
		level++
		calculateDeepestNode(node.left, level, deepestLevel, deepestNode)
		if level > *deepestLevel {
			deepestNode = node
			deepestLevel = &level
		}
		calculateDeepestNode(node.right, level, deepestLevel, deepestNode)
	}
}

// FindDeepest returns the deepest node if it was found
// TODO add tests
func (tree *SplayTree) FindDeepest() *Node {
	var level = 0
	var deepestLevel = new(int)
	var deepestNode *Node
	calculateDeepestNode(tree.root, level, deepestLevel, deepestNode)
	return deepestNode
}

// SplayDeepest finds the deepest node in the
// tree and splay it to root
func (tree *SplayTree) SplayDeepest() {
	// TODO implementation
}

// Get search for a node by its key and return it's value.
func (tree *SplayTree) Get(id int) {
	// TODO implementation
}

// GetNode search for a node by its key and return it
func (tree *SplayTree) GetNode(id int) {
	// TODO implementation
}

// Add adds a node to the tree and Splay it to root node.
func (tree *SplayTree) Add(id int) {
	// TODO implementation
}

// Remove removes a node from the tree
func (tree *SplayTree) Remove(id int) {
	// TODO implementation
}

// GetSize is a helper function that calculates tree size.
// TODO add tests
func (tree *SplayTree) GetSize() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.GetSize()
}

// GetSize is a helper function that calculates the size
// from a specific node
// TODO add tests
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
func Splay(node *Node, id int) {
	// TODO implementation
}

// RotateLeft is a helper that rotates node positions in a counterclockwise direction
// TODO add tests
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
// TODO add tests
func RotateRight(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	leftNode := node.left
	node.left = leftNode.left
	leftNode.right = node
	return leftNode
}
