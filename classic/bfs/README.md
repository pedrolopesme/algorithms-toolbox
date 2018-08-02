# Burrows-Wheeler transform


### Description

Breadth-first search (BFS) is an algorithm for traversing 
or searching tree or graph data structures. It starts at the 
tree root (or some arbitrary node of a graph, sometimes referred to as 
a 'search key'), and explores all of the neighbor nodes at 
the present depth prior to moving on to the nodes at the next depth level.

It uses the opposite strategy as depth-first search, which 
instead explores the highest-depth nodes first before being 
forced to backtrack and expand shallower nodes.

BFS and its application in finding connected 
components of graphs were invented in 1945 by Konrad Zuse 
and Michael Burke, in their (rejected) Ph.D. thesis on the 
Plankalkül programming language, but this was not published until 1972.
It was reinvented in 1959 by Edward F. Moore, who used it to 
find the shortest path out of a maze, and later developed by C. Y. 
Lee into a wire routing algorithm (published 1961).

Source: https://en.wikipedia.org/wiki/Breadth-first_search

![Animated BFS](https://upload.wikimedia.org/wikipedia/commons/4/46/Animated_BFS.gif)


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
