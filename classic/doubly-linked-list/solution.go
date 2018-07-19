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
	id       int
	date     time.Time
	next     *Node
	previous *Node
}

// TODO Implement and add Tests
func (list *LinkedList) Push(newNode *Node) {
	return
}

// TODO Implement and add Tests
func (list *LinkedList) Append(newNode *Node) {
	return
}


// TODO Implement and add Tests
func (list *LinkedList) InsertAfter(id int, newNode *Node) {
}

// TODO Implement and add Tests
func (list *LinkedList) Find(id int) *Node {
	return
}

// TODO Implement and add Tests
func (list *LinkedList) Remove(id int) {
	return
}
