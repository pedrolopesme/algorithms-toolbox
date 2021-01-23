package main

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) Prepend(n *Node) {
	if l.size == 0 {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.size++
}

func (l *List) Append(n *Node) {
	if l.size == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func main() {
	// silence is gold
}
