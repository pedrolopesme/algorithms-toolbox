# Stack


### Description

In computer science, a stack is an abstract data type that serves as a 
collection of elements, with two principal operations:

* push, which adds an element to the collection, and
* pop, which removes the most recently added element that was not yet removed.

The order in which elements come off a stack gives rise to its alternative name, 
LIFO (last in, first out). Additionally, a peek operation may give access to the 
top without modifying the stack. The name "stack" for this type of structure comes 
from the analogy to a set of physical items stacked on top of each other, which makes 
it easy to take an item off the top of the stack, while getting to an item deeper in 
the stack may require taking off multiple other items first.

Source: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)


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
