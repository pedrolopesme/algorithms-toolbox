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

// FindDeepest returns the deepest node if it was found
func (tree *SplayTree) FindDeepest() {
	// TODO implementation
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
func (tree *SplayTree) GetSize() {
	// TODO implementation
}

// Splay splays a node to the root of the tree. If there isn't a node with
// that key, the last node along the search path for the key will be splayed to
// the root.
func Splay(node *Node, id int) {
	// TODO implementation
}

// RotateLeft is a helper that rotates node positions in a counterclockwise direction
func (node *Node) RotateLeft() *Node {
	// TODO implementation
}

// RotateLeft is a helper that rotates node positions in a clockwise direction
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
