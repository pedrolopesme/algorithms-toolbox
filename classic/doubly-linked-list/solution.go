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
		list.last = newNode
		return
	}

	newNode.next = list.first
	list.first.previous = newNode
	list.first = newNode
}

// Append inserts a new node on the end of the list
func (list *LinkedList) Append(id int) {
	newNode := &Node{id: id}

	if list.last == nil {
		list.first = newNode
		list.last = newNode
		return
	}

	newNode.previous = list.last
	list.last.next = newNode
	list.last = newNode
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

// IsEmpty checks if the list has elements or not.
func (list *LinkedList) IsEmpty() bool {
	return list.first == nil
}
