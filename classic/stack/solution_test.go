package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPushToAnEmptyStack(test *testing.T) {
	stack := Stack{}
	stack.Push(&Node{id: 1})
	assert.False(test, stack.IsEmpty())
	assert.Equal(test, 1, stack.first.id)
}

func TestPushToANonEmptyStack(test *testing.T) {
	stack := Stack{}
	stack.Push(&Node{id: 1})
	stack.Push(&Node{id: 2})
	stack.Push(&Node{id: 3})
	assert.False(test, stack.IsEmpty())
	assert.Equal(test, 3, stack.first.id)
	assert.Equal(test, 2, stack.first.next.id)
	assert.Equal(test, 1, stack.first.next.next.id)
}

func TestPopFromAnEmptyStack(test *testing.T) {
	stack := Stack{}
	assert.Nil(test, stack.Pop())
	assert.True(test, stack.IsEmpty())
}

func TestPopFromANonEmptyStack(test *testing.T) {
	stack := Stack{}
	stack.Push(&Node{id: 1})
	stack.Push(&Node{id: 2})
	stack.Push(&Node{id: 3})
	assert.False(test, stack.IsEmpty())

	assert.Equal(test, 3, stack.Pop().id)
	assert.False(test, stack.IsEmpty())

	assert.Equal(test, 2, stack.Pop().id)
	assert.False(test, stack.IsEmpty())

	assert.Equal(test, 1, stack.Pop().id)
	assert.True(test, stack.IsEmpty())
}


func TestIsEmptyAtAnEmptyStack(test *testing.T) {
	stack := Stack{}
	assert.True(test, stack.IsEmpty())
}

func TestIsEmptyAtANonEmptyStack(test *testing.T) {
	stack := Stack{first: &Node{}}
	assert.False(test, stack.IsEmpty())
}
