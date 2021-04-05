package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMergeSort_BaseCases(t *testing.T) {
	l := List{}
	l.Add(&Node{val: 1})
	l.Add(&Node{val: 2})
	l.Add(&Node{val: 3})
	l.Add(&Node{val: 4})
	l.Add(&Node{val: 5})
	assert.Equal(t, 5, l.head.val)
	assert.Equal(t, 4, l.head.next.val)
	assert.Equal(t, 3, l.head.next.next.val)
	assert.Equal(t, 2, l.head.next.next.next.val)
	assert.Equal(t, 1, l.head.next.next.next.next.val)

	sortedL := MergeSort(l.head)
	assert.Equal(t, 1, sortedL.val)
	assert.Equal(t, 2, sortedL.next.val)
	assert.Equal(t, 3, sortedL.next.next.val)
	assert.Equal(t, 4, sortedL.next.next.next.val)
	assert.Equal(t, 5, sortedL.next.next.next.next.val)
}

func TestGetMiddle_WhenPassAnLinkedList_ThenReturnMiddlePosition(t *testing.T) {
	l := List{}
	l.Add(&Node{val: 1})
	l.Add(&Node{val: 2})
	l.Add(&Node{val: 3})
	l.Add(&Node{val: 4})
	l.Add(&Node{val: 5})

	middle := GetMiddle(l.head)
	assert.Equal(t, 3, middle.val)
}
