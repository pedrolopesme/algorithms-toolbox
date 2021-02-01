package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMakeAnagram_WhenSmiliarStringsAreGiven_ThenThereIsNoChanges(t *testing.T) {
	assert.Equal(t, int32(0), makeAnagram("abc", "cba"))
}

func TestMakeAnagram_WhenDifferentStringsWereGiven_ThenThereIsNoChanges(t *testing.T) {
	assert.Equal(t, int32(4), makeAnagram("cde", "abc"))
}
