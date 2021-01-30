# Linked List


### Description

In computer science, a linked list is a linear collection of data elements, 
whose order is not given by their physical placement in memory. 
Instead, each element points to the next. It is a data structure consisting of a 
collection of nodes which together represent a sequence. In its most basic form, 
each node contains: data, and a reference (in other words, a link) to the next node 
in the sequence. This structure allows for efficient insertion or removal of elements
 from any position in the sequence during iteration.

Source: https://en.wikipedia.org/wiki/Linked_list

### Complexity

#### Time

Average:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| Θ(n) | Θ(n) | Θ(1) | Θ(1) |

Worst:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| O(n) | O(n) | O(n) | O(n) |


#### Space

Worst: O(n)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
