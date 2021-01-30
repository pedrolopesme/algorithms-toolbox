package main

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"time"
)

func TestAddNodeToAEmptyList(test *testing.T) {
	node := &Node{id: 1, date: time.Now()}
	list := LinkedList{}
	list.Add(node)
	assert.Equal(test, uint64(1), list.size)
	assert.Equal(test, node.id, list.first.id)
	assert.Equal(test, node.id, list.last.id)
}

func TestAddNodeToANonEmptyList(test *testing.T) {
	node1 := &Node{id: 1, date: time.Now()}
	node2 := &Node{id: 2, date: time.Now()}
	node3 := &Node{id: 3, date: time.Now()}
	list := LinkedList{}
	list.Add(node1)
	list.Add(node2)
	list.Add(node3)
	assert.Equal(test, uint64(3), list.size)
	assert.Equal(test, node1.id, list.first.id)
	assert.Equal(test, node3.id, list.last.id)
}

func TestRemoveNodeFromAEmptyList(test *testing.T) {
	list := LinkedList{}
	list.Remove(1)
	assert.Equal(test, uint64(0), list.size)
}

func TestRemoveAnExistentNodeFromAListWithOneNode(test *testing.T) {
	node := &Node{id: 1, date: time.Now()}
	list := LinkedList{}
	list.Add(node)
	list.Remove(1)
	assert.Equal(test, uint64(0), list.size)
}

func TestRemoveAnExistentNodeFromAListWithMultipleNodes(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})
	list.Remove(2)
	assert.Equal(test, uint64(2), list.size)
	assert.Equal(test, 1, list.first.id)
	assert.Equal(test, 3, list.last.id)
}

func TestRemoveAnExistentNodeFromTheBeginningOfAList(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})
	list.Remove(1)
	assert.Equal(test, uint64(2), list.size)
	assert.Equal(test, 2, list.first.id)
	assert.Equal(test, 3, list.last.id)
}

func TestRemoveAnExistentNodeFromTheEndOfAList(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})
	list.Remove(3)
	assert.Equal(test, uint64(2), list.size)
	assert.Equal(test, 1, list.first.id)
	assert.Equal(test, 2, list.last.id)
}

func TestRemoveANonExistentNode(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})
	list.Remove(4)
	assert.Equal(test, uint64(3), list.size)
	assert.Equal(test, 1, list.first.id)
	assert.Equal(test, 3, list.last.id)
}

func TestFindAnExistentNode(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})

	assert.Equal(test, 3, list.Find(3).id)
}

func TestFindANonExistentNode(test *testing.T) {
	list := LinkedList{}
	list.Add(&Node{id: 1, date: time.Now()})
	list.Add(&Node{id: 2, date: time.Now()})
	list.Add(&Node{id: 3, date: time.Now()})

	assert.Nil(test, list.Find(4))
}