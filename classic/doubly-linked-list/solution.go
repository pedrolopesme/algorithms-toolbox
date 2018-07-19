package main

type LinkedList struct {
	size  uint64
	first *Node
	last  *Node
}

type Node struct {
	id       int
	next     *Node
	previous *Node
}

// Push inserts a new node on the front of the list
func (list *LinkedList) Push(id int) {
	newNode := &Node{id: id}

	if list.first == nil {
		list.first = newNode
		return
	}

	newNode.next = list.first
	list.first.previous = newNode
	list.first = newNode
}

// TODO Implement and add Tests
func (list *LinkedList) Append(id int) {
	return
}

// TODO Implement and add Tests
func (list *LinkedList) InsertAfter(id int, newEntry int) {
}

// TODO Implement and add Tests
func (list *LinkedList) Find(id int) *Node {
	return nil
}

// TODO Implement and add Tests
func (list *LinkedList) Remove(id int) {
	return
}

// TODO Implement and add Tests
func (list *LinkedList) IsEmpty() bool {
	return list.first == nil
}
