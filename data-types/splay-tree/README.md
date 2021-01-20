# Splay Tree


### Description

A splay tree is a self-adjusting binary search tree with 
the additional property that recently accessed elements are 
quick to access again. It performs basic operations such as insertion, 
look-up and removal in O(log n) amortized time. For many sequences of 
non-random operations, splay trees perform better than other search trees, 
even when the specific pattern of the sequence is unknown. The splay tree was 
invented by Daniel Sleator and Robert Tarjan in 1985.

Source: https://en.wikipedia.org/wiki/Splay_tree


### Complexity

#### Time

Average:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| - | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) |

Worst:

| Access | Search | Insertion | Deletion | 
|---|---|---|---|
| - | O(log(n)) | O(log(n)) | O(log(n)) |


#### Space

Worst: O(n)



### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
