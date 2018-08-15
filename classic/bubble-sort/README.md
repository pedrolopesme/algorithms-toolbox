# Bubble Sort


### Description

Bubble sort, sometimes referred to as sinking sort, 
is a simple sorting algorithm that repeatedly steps 
through the list to be sorted, compares each pair of 
adjacent items and swaps them if they are in the wrong order. 
The pass through the list is repeated until no swaps are 
needed, which indicates that the list is sorted. 

The algorithm, which is a comparison sort, is 
named for the way smaller or larger elements "bubble" to the top 
of the list. Although the algorithm is simple, it is 
too slow and impractical for most problems even when compared 
to insertion sort. Bubble sort can be practical if the input 
is in mostly sorted order with some out-of-order elements nearly in position.

Source: https://en.wikipedia.org/wiki/Bubble_sort


### Complexity


#### Time

Average:
O(n^2)

Worst:
O(n^2)

#### Space

Worst: O(1)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
