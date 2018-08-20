package main

// TableSize defines the buckets quantity in my hash table
const TableSize = 256

// HashEntry represents an entry with an int key
// and an empty interface to hold the value
// TODO add new item pointer in order to create a singly linked list
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
	table map[int]HashEntry
}

// NewHashTable knows how to build a HashTable
func NewHashTable() HashTable {
	return HashTable{
		size:  TableSize,
		table: make(map[int]HashEntry),
	}
}

// Get returns the hashentry from a given key
// TODO what if occurs a hash collision?
func (ht HashTable) Get(key int) (entry HashEntry) {
	hash := hashFunc(key)
	if val, ok := ht.table[hash]; ok {
		entry = val
	}
	return
}

// Put adds a hash entry into my table
// TODO what if occurs a hash collision?
func (ht HashTable) Put(entry HashEntry) {
	hash := hashFunc(entry.key)
	ht.table[hash] = entry
}

// hashFunc knows how to generate a hash from key
func hashFunc(key int) int {
	return key % TableSize
}
