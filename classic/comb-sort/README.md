# Comb sort


### Description

Comb sort is a relatively simple sorting algorithm 
based on bubble sort and originally designed by Wlodzimierz Dobosiewicz 
in 1980. It was later rediscovered and popularized by 
Stephen Lacey and Richard Box with a Byte Magazine 
article published in April 1991. The basic idea is to 
eliminate turtles, or small values near the end of the list, 
since in a bubble sort these slow the sorting down tremendously. 
(Rabbits, large values around the beginning of the list, 
do not pose a problem in bubble sort) It accomplishes 
this by initially swapping elements that are a certain distance 
from one another in the array, rather than only swapping elements if they are 
adjacent to one another, and then shrinking the chosen distance until it is 
operating as a normal bubble sort. 

![Comb sort](https://upload.wikimedia.org/wikipedia/commons/4/46/Comb_sort_demo.gif)

Source: https://en.wikipedia.org/wiki/Comb_sort


### Complexity


#### Time

Average:
Omega(n^2/2^p), where p is the number of increments


Worst:
O(n^2)

#### Space

Worst: O(1)


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* test: runs all tests
* fmt: runs gofmt for all go files
