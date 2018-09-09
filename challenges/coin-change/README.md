# Coin Change Problem

### Description
Given a value N, if we want to make change for N cents, and we have 
infinite supply of each of S = {1,5,10,25,50} valued coins, 
what is the minimum number of coins to make the change?

### Input
You'll receive a float value V, where 0 <= V <= 2^32 

### Output
Ex:  The number of coins need for the change, followed by 
a list of its values:

total of coins: 5 
2 coins of 50
1 coins of 25
2 coins of 1


### Makefile
This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* build: generates binaries
* test: runs all tests
* clean: removes binaries
* run: executes main func
* fmt: runs gofmt for all go files
