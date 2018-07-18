package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsEmptyAtAnEmptyStack(test *testing.T) {
	stack := Stack{}
	assert.True(test, stack.IsEmpty())
}

func TestIsEmptyAtANonEmptyStack(test *testing.T) {
	stack := Stack{ first: &Node{} }
	assert.False(test, stack.IsEmpty())
}
