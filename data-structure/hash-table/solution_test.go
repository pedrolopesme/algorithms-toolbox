package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPutEntryToAEmptyHashTable(test *testing.T) {
	entry := NewEntry(1, "Foo")
	table := NewHashTable()
	table.Put(entry)
	assert.Equal(test, entry, table.Get(1))
}


func TestPutMultipleEntriesToAHashTable(test *testing.T) {
	table := NewHashTable()
	table.Put(NewEntry(1, "Foo"))
	table.Put(NewEntry(2, "Bar"))
	table.Put(NewEntry(3, "Baz"))
	assert.Equal(test, 1, table.Get(1).key)
	assert.Equal(test, 2, table.Get(2).key)
	assert.Equal(test, 3, table.Get(3).key)
}

func TestGetAnInvalidEntryFromAHashTable(test *testing.T) {
	table := NewHashTable()
	table.Put(NewEntry(1, "Foo"))
	assert.Empty(test,table.Get(10))
}