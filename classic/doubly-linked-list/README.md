# Linked List


### Description

In computer science, a doubly linked list is a linked data structure that 
consists of a set of sequentially linked records called nodes. 
Each node contains two fields, called links, that are references to 
the previous and to the next node in the sequence of nodes. The 
beginning and ending nodes' previous and next links, respectively, 
point to some kind of terminator, typically a sentinel node or null, 
to facilitate traversal of the list. If there is only one sentinel node, 
then the list is circularly linked via the sentinel node. It can be conceptualized 
as two singly linked lists formed from the same data items, but in 
opposite sequential orders.

![Doubly Linked Image](https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Doubly-linked-list.svg/610px-Doubly-linked-list.svg.png)

Source: https://en.wikipedia.org/wiki/Doubly_linked_list

### Complexity

#### Time

Average:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| Θ(n) | Θ(n) | Θ(1) | Θ(1) |

Worst:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| O(n) | O(n) | O(1) | O(1) |


#### Space

Worst: O(n)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
