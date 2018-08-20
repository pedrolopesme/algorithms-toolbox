package main

// TableSize defines the buckets quantity in my hash table
const TableSize = 256

// HashEntry represents an entry with an int key
// and an empty interface to hold the value
type HashEntry struct {
	key   int
	value interface{}
}

// NewEntry knows how to build a HashEntry for a given
// key/value pair.
func NewEntry(key int, value interface{}) HashEntry {
	return HashEntry{
		key:   key,
		value: value,
	}
}

// IHashTable contains basic HashTable operations
type IHashTable interface {
	Get(key int) HashEntry
	Put(entry HashEntry)
}

// HashTable
type HashTable struct {
	size  int
	table []HashEntry
}

// NewHashTable knows how to build a HashTable
func NewHashTable() HashTable {
	return HashTable{
		size:  TableSize,
		table: make([]HashEntry, TableSize, TableSize),
	}
}

// Get returns the hashentry from a given key
// TODO implement and add tests
func (ht HashTable) Get(key int) HashEntry {
	return HashEntry{}
}

// Put adds a hash entry into my table
// TODO implement and add tests
func (ht HashTable) Put(entry HashEntry) {
}
