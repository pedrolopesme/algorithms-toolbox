package main

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	size int
}

func (l *List) Prepend(n *Node) {
	if l.size == 0 {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.size++
}

func (l *List) Append(n *Node) {
	if l.size == 0 {
		l.head = n
	} else {
		iNode := l.head
		for iNode.next != nil {
			iNode = iNode.next
		}
		iNode.next = n
	}
	l.size++
}

func main() {
	// silence is gold
}
