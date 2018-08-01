# Two String Common Substrings


### Description

Given two strings, determine if they share a common substring. A substring may be as small as one character.

For example, the words "a", "and", "art" share the common substring . The words "be" and "cat" do not share a substring.

##### Function Description

Complete the function twoStrings in the editor below. It should return a string, either YES or NO based on whether the strings share a common substring.

twoStrings has the following parameter(s):

* s1, s2: two strings to analyze .

##### Input Format

The first line contains a single integer , the number of test cases.

The following  pairs of lines are as follows:

* The first line contains string .
* The second line contains string .


##### Constraints

* s1 and s2 consist of characters in the range ascii[a-z].
* p is greater or equal to 1 and lower or equal to 10
* s1 and s2 is greater or equal to 1 and lower or equal to 10^5 

##### Output Format

For each pair of strings, return YES or NO.

Sample Input:

```
2
hello
world
hi
world
```

##### Sample Output
```
YES
NO
```

#####  Explanation

We have p pairs to check:

* s1 = "hello", s2 = "world". The substrings "o" and "l" are common to both strings.
* a = "h1", b = "world", s1 and s2 share no common substrings.


### Makefile

This project provides a Makefile with all common operations need to develop, 
test and build the binary.

* build: generates binaries
* test: runs all tests
* clean: removes binaries
* run: executes main func
* fmt: runs gofmt for all go files
