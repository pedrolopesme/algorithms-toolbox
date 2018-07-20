package main

type DoublyLinkedList struct {
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
func (list *DoublyLinkedList) Push(id int) {
	newNode := &Node{id: id}
	list.size++

	if list.first == nil {
		list.first = newNode

		// list has only the new node
		if list.last == nil {
			list.last = newNode
		}
		return
	}

	newNode.next = list.first
	list.first.previous = newNode
	list.first = newNode
}

// Append inserts a new node on the end of the list
func (list *DoublyLinkedList) Append(id int) {
	newNode := &Node{id: id}
	list.size++

	if list.last == nil {
		list.last = newNode

		// list has only the new node
		if list.first == nil {
			list.first = newNode
		}
		return
	}

	newNode.previous = list.last
	list.last.next = newNode
	list.last = newNode
}

// TODO Implement and add Tests
func (list *DoublyLinkedList) InsertAfter(id int, newEntry int) {
}

func (list *DoublyLinkedList) Find(id int) *Node {
	if list.size == 0 {
		return nil
	}

	currentNode := list.first
	for currentNode.id != id {
		if currentNode.next == nil {
			return nil
		}
		currentNode = currentNode.next
	}
	return currentNode
}

// TODO Implement and add Tests
func (list *DoublyLinkedList) Remove(id int) {
	return
}

// IsEmpty checks if the list has elements or not.
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.first == nil
}
