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

func (l *List) Insert(pos int, n *Node) {
	// Dealing with insertion into the head
	if pos == 0 {
		l.Prepend(n)
		return
	}

	// Dealing with insertion into the tail
	if l.size <= pos {
		l.Append(n)
		return
	}

	// Dealing with insertion in the middle
	iPos := 0
	iNode := l.head
	var prevNode *Node
	for iPos != pos {
		iPos++
		prevNode = iNode
		iNode = iNode.next
	}
	prevNode.next = n
	n.next = iNode
	l.size++
}

func main() {
	// silence is gold
}
