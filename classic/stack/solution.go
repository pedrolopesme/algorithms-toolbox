package main

// Node containing data and a pointer to next node
type Node struct {
	id   int
	next *Node
}

// Stack defines the basic tree structure
type Stack struct {
	root *Node
}

// Push add a node to beginning of the Stack
func (s *Stack) Push(node *Node) {
	// TODO implement and test
}

// Pop remove a node of beginning of the Stack
func (s *Stack) Pop(node *Node) {
	// TODO implement and test
}

// Peek return the reference from the Node at the top
// of the stack
func (s *Stack) Peek(node *Node) {
	// TODO implement and test
}

// IsEmpty confirms if the stack is empty
func (s *Stack) IsEmpty(node *Node) {
	// TODO implement and test
}
