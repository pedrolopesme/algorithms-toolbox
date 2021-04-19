package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMagazine_Case1(t *testing.T) {
	magazine := []string{"attack", "at", "dawn"}
	note := []string{"Attack", "at", "dawn"}
	assert.False(t, checkMagazine(magazine, note))
}

func TestCheckMagazine_Case2(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	assert.True(t, checkMagazine(magazine, note))
}

func TestCheckMagazine_Case3(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}
	assert.False(t, checkMagazine(magazine, note))
}

func TestCheckMagazine_Case4(t *testing.T) {
	magazine := []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
	note := []string{"ive", "got", "some", "coconuts"}
	assert.False(t, checkMagazine(magazine, note))
}
