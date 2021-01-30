# Insertion Sort


### Description

In computer science, merge sort (also commonly spelled mergesort) 
is an efficient, general-purpose, comparison-based sorting algorithm. 
Most implementations produce a stable sort, which means that the implementation 
preserves the input order of equal elements in the sorted output. Merge sort 
is a divide and conquer algorithm that was invented by John von Neumann in 1945.

A detailed description and analysis of bottom-up mergesort appeared in a 
report by Goldstine and von Neumann as early as 1948.

![Merge Sort](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

Source: https://en.wikipedia.org/wiki/Merge_sort


### Complexity


#### Time

Average:
Θ(n log(n))


Worst:
O(n log(n))

#### Space

Worst: O(n)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
