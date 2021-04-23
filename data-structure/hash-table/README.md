# Hash Table

### Description

In computing, a hash table (hash map) is a data structure 
that implements an associative array abstract data type, a structure 
that can map keys to values. A hash table uses a hash function to compute 
an index into an array of buckets or slots, from which the desired value can be found.

Ideally, the hash function will assign each key to a unique bucket, 
but most hash table designs employ an imperfect hash function, 
which might cause hash collisions where the hash function generates the same 
index for more than one key. Such collisions must be accommodated in some way.

In a well-dimensioned hash table, the average cost (number of instructions) 
for each lookup is independent of the number of elements stored in the table. 
Many hash table designs also allow arbitrary insertions and deletions 
of key-value pairs, at constant average cost per operation.

In many situations, hash tables turn out to be on average more efficient 
than search trees or any other table lookup structure. For this reason, 
they are widely used in many kinds of computer software, particularly for 
associative arrays, database indexing, caches, and sets.

Source: https://en.wikipedia.org/wiki/Hash_table

![Hash Table](https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/Hash_table_3_1_1_0_1_0_0_SP.svg/315px-Hash_table_3_1_1_0_1_0_0_SP.svg.png)

### Complexity

#### Time
Average:
Θ(1)

Worst:
O(n)

#### Space
Worst: O(n)


### Makefile
This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
