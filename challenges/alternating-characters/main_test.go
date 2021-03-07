package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestAlternatingCharacters_BaseCases(t *testing.T) {
	assert.Equal(t, int32(3), alternatingCharacters("AAAA"))
	assert.Equal(t, int32(4), alternatingCharacters("BBBBB"))
	assert.Equal(t, int32(0), alternatingCharacters("ABABABAB"))
	assert.Equal(t, int32(0), alternatingCharacters("BABABA"))
	assert.Equal(t, int32(4), alternatingCharacters("AAABBB"))
}
