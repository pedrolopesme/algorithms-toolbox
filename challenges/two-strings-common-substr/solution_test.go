package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoStringsWithSubstring(test *testing.T) {
	str1 := "Hello World"
	str2 := "Hey Jude"
	assert.Equal(test, "YES", twoStrings(str1, str2))
}


func TestTwoStringsWithotSubstring(test *testing.T) {
	str1 := "Hello,World"
	str2 := "My big cat"
	assert.Equal(test, "NO", twoStrings(str1, str2))
}
