# Heap Sort


### Description

In computer science, heapsort is a comparison-based 
sorting algorithm. Heapsort can be thought of as an 
improved selection sort: like that algorithm, it divides 
its input into a sorted and an unsorted region, and it 
iteratively shrinks the unsorted region by extracting the 
largest element and moving that to the sorted region. The 
improvement consists of the use of a heap data structure 
rather than a linear-time search to find the maximum.

Although somewhat slower in practice on most machines than 
a well-implemented quicksort, it has the advantage of a more 
favorable worst-case O(n log n) runtime. Heapsort is an in-place 
algorithm, but it is not a stable sort.

Source: https://en.wikipedia.org/wiki/Heapsort

![Insertion Sort](https://upload.wikimedia.org/wikipedia/commons/1/1b/Sorting_heapsort_anim.gif)

### Complexity


#### Time

Average:
Θ(n log(n))

Worst:
O(n log(n))

#### Space

Worst: O(1)



### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
