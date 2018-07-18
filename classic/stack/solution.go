package main

// Node containing data and a pointer to next node
type Node struct {
	id   int
	next *Node
}

// StackQueue defines the basic tree structure
type StackQueue struct {
	root *Node
}

// Push add a node to beginning of the Stack
func (s *StackQueue) Push(node *Node) {
	// TODO implement and test
}

// Pop remove a node of beginning of the Stack
func (s *StackQueue) Pop(node *Node) {
	// TODO implement and test
}

// Peek return the reference from the Node at the top
// of the stack
func (s *StackQueue) Peek(node *Node) {
	// TODO implement and test
}

// IsEmpty confirms if the stack is empty
func (s *StackQueue) IsEmpty(node *Node) {
	// TODO implement and test
}
