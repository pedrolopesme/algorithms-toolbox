# Count sort


### Description

In computer science, counting sort is an algorithm for 
sorting a collection of objects according to keys that are small 
integers; that is, it is an integer sorting algorithm. 

It operates by counting the number of objects that have 
each distinct key value, and using arithmetic on those counts 
to determine the positions of each key value in the output sequence. 
Its running time is linear in the number of items and the difference
between the maximum and minimum key values, so it is only suitable for 
direct use in situations where the variation in keys is not significantly 
greater than the number of items. However, it is often used as a subroutine 
in another sorting algorithm, radix sort, that can handle larger keys more efficiently

Because counting sort uses key values as indexes into an array, 
it is not a comparison sort, and the Ω(n log n) lower bound for 
comparison sorting does not apply to it.
 
Bucket sort may be used for many of the same tasks as counting sort, 
with a similar time analysis; however, compared to counting sort, bucket 
sort requires linked lists, dynamic arrays or a large amount of preallocated 
memory to hold the sets of items within each bucket, whereas counting sort 
instead stores a single number (the count of items) per bucket.

Source: https://en.wikipedia.org/wiki/Counting_sort


### Complexity


#### Time

Average:
O(n + k), where n is the number of elements, and k is the number of keys

#### Space
Worst: O(k)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
