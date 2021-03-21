package main

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Add(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
}

func GetMiddle(n *Node) *Node {
	if n == nil {
		return n
	}

	fast := n
	slow := n
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func Merge(l, r *Node) *Node {
	if l == nil {
		return r
	} else if r == nil {
		return l
	}

	var result *Node
	if l.val <= r.val {
		result = l
		result.next = Merge(l.next, r)
	} else {
		result = r
		result.next = Merge(l, r.next)
	}
	return result
}

func MergeSort(n *Node) *Node {
	if n == nil || n.next == nil {
		return n
	}

	middle := GetMiddle(n)
	nextMiddle := middle.next

	// sorting the first part of the list
	middle.next = nil
	left := MergeSort(n)

	right := MergeSort(nextMiddle)
	return Merge(left, right)
}

func main() {
}
