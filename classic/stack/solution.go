package main

// Node containing data and a pointer to next node
type Node struct {
	id   int
	next *Node
}

// Stack defines the basic tree structure
type Stack struct {
	first *Node
}

// Push add a node to beginning of the Stack
func (s *Stack) Push(newNode *Node) {
	if s.first == nil {
		s.first = newNode
		return
	}

	newNode.next = s.first
	s.first = newNode
}

// Pop remove a node of beginning of the Stack
func (s *Stack) Pop() (node *Node) {
	if s.first == nil {
		return
	}

	poppedNode := s.first
	s.first = s.first.next
	return poppedNode
}

// Peek return the reference from the Node at the top
// of the stack
func (s *Stack) Peek() (node *Node) {
	// TODO implement and test
	return
}

// IsEmpty confirms if the stack is empty
func (s *Stack) IsEmpty() bool {
	if s.first == nil {
		return true
	}
	return false
}
