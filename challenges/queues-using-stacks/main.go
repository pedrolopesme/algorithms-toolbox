package main

type Node struct {
	val  int
	next *Node
	prev *Node
}

type Stack struct {
	head *Node
	tail *Node
	size int
}

func (s *Stack) Push(n *Node) {
	if s.size == 0 {
		s.head = n
		s.tail = n
	} else {
		n.prev = s.tail
		s.tail.next = n
		s.tail = n
	}

	s.size++
}

func (s *Stack) Pop() *Node {
	if s.size == 0 {
		return nil
	}

	var node *Node
	if s.size == 1 {
		node = s.head
		s.head = nil
		s.tail = nil
	} else {
		node = s.tail
		s.tail = node.prev
		node.prev = nil
	}
	s.size--
	return node
}

func main() {
	//silence is gold
}
