# Radix sort


### Description

In computer science, radix sort is a non-comparative integer sorting algorithm 
that sorts data with integer keys by grouping keys by the individual digits 
which share the same significant position and value. 
A positional notation is required, but because integers can represent strings of characters 
(e.g., names or dates) and specially formatted floating point numbers, 
radix sort is not limited to integers. Radix sort dates back as far as 1887 
to the work of Herman Hollerith on tabulating machines.

Most digital computers internally represent all of their 
data as electronic representations of binary numbers, so processing the digits 
of integer representations by groups of binary digit representations is 
most convenient. Radix sorts can be implemented to start at either 
the most significant digit (MSD) or least significant digit (LSD). 
For example, when sorting the number 1234 into a list, one could start with the 1 or the 4

Source: https://en.wikipedia.org/wiki/Radix_sort

### Complexity

#### Time
Average:
Θ(nk), for k keys which are integers of word size n.

Worst:
O(nk), , for k keys which are integers of word size n.

#### Space
Worst: O(n+k)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
