# Depth-first search


### Description

Depth-first search (DFS) is an algorithm for traversing 
or searching tree or graph data structures. The algorithm 
starts at the root node (selecting some arbitrary node as the 
root node in the case of a graph) and explores as far as possible 
along each branch before backtracking.

A version of depth-first search was investigated in the 19th 
century by French mathematician Charles Pierre Trémaux
 as a strategy for solving mazes.

Source: https://en.wikipedia.org/wiki/Depth-first_search


### Complexity

##### Time

O(|V|+|E|), where |V| is the number of vertices and |E|, edges.

##### Space

O(|V|), where |V| is the cardinality of the set of the vertices.


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
