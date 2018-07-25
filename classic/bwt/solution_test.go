package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutations(test *testing.T) {
	input := "TOAD"
	expectedOutput := []string{
		"TOAD",
		"OADT",
		"ADTO",
		"DTOA",
	}
	assert.Equal(test, expectedOutput, generatePermutations(input))
}

func TestPermutationsWithEmptyString(test *testing.T) {
	input := ""
	var expectedOutput []string
	assert.Equal(test, expectedOutput, generatePermutations(input))
}

func TestStringTransformation(test *testing.T) {
	input := "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES"
	expectedOutput := "TEXYDST.E.IXIXIXXSSMPPS.B..E.S.EUSFXDIIOIIIT"
	assert.Equal(test, expectedOutput, Transform(input))
}

func TestStringTransformationWithNoString(test *testing.T) {
	input := ""
	expectedOutput := ""
	assert.Equal(test, expectedOutput, Transform(input))
}