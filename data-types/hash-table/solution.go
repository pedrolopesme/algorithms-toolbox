package main

// TableSize defines the buckets quantity in my hash table
const TableSize = 256

// HashEntry represents an entry with an int key
// and an empty interface to hold the value
type HashEntry struct {
	key   int
	value interface{}
	next  *HashEntry
}

// NewEntry knows how to build a HashEntry for a given
// key/value pair.
func NewEntry(key int, value interface{}) HashEntry {
	return HashEntry{
		key:   key,
		value: value,
	}
}

// Append add a entry to the of HashEntries singly linked list
func (he HashEntry) Append(entry *HashEntry) {
	var node *HashEntry
	if he.next == nil {
		node = &he
	} else {
		node = he.next
		for node.next != nil {
			node = node.next
		}
	}
	node.next = entry
}

// Search iterates over Entries until find the given key
func (he HashEntry) Search(key int) (entry *HashEntry) {
	if he.key == key {
		entry = &he
	} else {
		node := he.next
		for node != nil && node.key != key && node.next != nil {
			node = node.next
		}
	}
	return
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

// Get returns the hash entry from a given key
func (ht HashTable) Get(key int) (entry HashEntry) {
	hash := hashFunc(key)
	if val, ok := ht.table[hash]; ok {
		entry = *val.Search(key)
	}
	return
}

// Put adds a hash entry into my table
func (ht HashTable) Put(entry HashEntry) {
	hash := hashFunc(entry.key)
	if val, ok := ht.table[hash]; ok {
		val.Append(&entry)
	} else {
		ht.table[hash] = entry
	}
}

// hashFunc knows how to generate a hash from key
func hashFunc(key int) int {
	return key % TableSize
}
