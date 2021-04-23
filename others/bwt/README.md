# Burrows-Wheeler transform


### Description

The Burrows–Wheeler transform (BWT, also called block-sorting compression) 
rearranges a character string into runs of similar characters. 
This is useful for compression, since it tends to be easy to compress a string 
that has runs of repeated characters by techniques such as move-to-front 
transform and run-length encoding. More importantly, the transformation is 
reversible, without needing to store any additional data. The BWT is thus a 
"free" method of improving the efficiency of text compression algorithms, 
costing only some extra computation.

The Burrows–Wheeler transform is an algorithm used to prepare data 
for use with data compression techniques such as bzip2. It was invented by 
Michael Burrows and David Wheeler in 1994 while Burrows was working at DEC 
Systems Research Center in Palo Alto, California. It is based on a previously 
unpublished transformation discovered by Wheeler in 1983. The algorithm can be 
implemented efficiently using a suffix array thus reaching linear time complexity.

Source: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform

### Complexity

#### Time

Worst:

O(n+σ) time, where n is the length of string and σ is the length of the alphabet.


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
