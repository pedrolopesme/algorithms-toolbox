package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReversal_WhenStringIsGiven_ReturnsTheInvertedString(t *testing.T) {
	given := "Pedro"
	expected := "ordeP"
	assert.Equal(t, expected, StringReversal(given))
}

func TestStringReversal_WhenEmptyStringIsGiven_ReturnsTheInvertedString(t *testing.T) {
	given := ""
	expected := ""
	assert.Equal(t, expected, StringReversal(given))
}
