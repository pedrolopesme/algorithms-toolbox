package main

import (
	"time"
)

type LinkedList struct {
	size  uint64
	first *Node
	last  *Node
}

type Node struct {
	id   int
	date time.Time
	next *Node
}

func (list *LinkedList) Add(newNode *Node) {
	if list.size == 0 {
		list.first = newNode
		list.last = newNode
	} else {
		lastNode := list.last
		lastNode.next = newNode
		list.last = newNode
	}
	list.size++
}

func (list *LinkedList) Find(id int) *Node {
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

func (list *LinkedList) Remove(id int) {
	if list.size == 0 {
		return
	}

	var previousNode *Node
	currentNode := list.first

	for currentNode.id != id {
		if currentNode.next == nil {
			return
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode.id == id {
		if previousNode == nil {
			// What if the Node was the first?
			currentNode = currentNode.next
			list.first = currentNode
		} else if currentNode.next == nil {
			// What if the Node was the last?
			previousNode.next = nil
			list.last = previousNode
		} else {
			// What if the Node was in between first and last?
			previousNode.next = currentNode.next
		}
		list.size--
	}
}
