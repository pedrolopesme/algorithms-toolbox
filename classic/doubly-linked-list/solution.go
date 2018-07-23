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

// InsertAfter inserts a new node after a given previous Id. If the
// previous ID does'nt exist, it tries to use the closest ID lower
// than the given one.
func (list *DoublyLinkedList) InsertAfter(previousId int, newEntry int) {
	if list.IsEmpty() {
		list.Append(newEntry)
		return
	}

	if list.first.id == previousId {
		list.Append(newEntry)
		return
	}

	if list.last.id == previousId {
		list.Append(newEntry)
		return
	}

	var previousNode *Node
	nodeIterator := list.first
	for nodeIterator != nil {
		// its a match. Use this node as previousNode.
		if nodeIterator.id == previousId {
			previousNode = nodeIterator
			break

			// haven't found a match yet and no node was chose. Using
			// the first node lower than the new entry
		} else if previousNode == nil && nodeIterator.id < newEntry {
			previousNode = nodeIterator

			// haven't found a match yet but the current node iterator
			// is closer to the newEntry than the current previousNode.
		} else if previousNode != nil && nodeIterator.id < newEntry && previousNode.id < nodeIterator.id {
			previousNode = nodeIterator
		}

		nodeIterator = nodeIterator.next
	}

	// No previous was found or it is the last node,
	// just append id to the list
	if previousNode == nil || previousNode.next == nil {
		list.Append(newEntry)
	} else {
		newNode := &Node{id: newEntry}
		newNode.previous = previousNode
		newNode.next = previousNode.next
		previousNode.next.previous = newNode
		previousNode.next = newNode
	}
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
