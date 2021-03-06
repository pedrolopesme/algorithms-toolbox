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
	_, tranformation := Transform(input)
	assert.Equal(test, expectedOutput, tranformation)
}

func TestStringTransformationWithNoString(test *testing.T) {
	input := ""
	expectedOutput := ""
	_, tranformation := Transform(input)
	assert.Equal(test, expectedOutput, tranformation)
}

func TestStringInverseTransformation(test *testing.T) {
	expectedOutput := "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES"
	index, transformation := Transform(expectedOutput)
	assert.Equal(test, expectedOutput, InverseTransform(transformation, index))
}

func TestStringInverseTransformationWithEmptyString(test *testing.T) {
	expectedOutput := ""
	index, transformation := Transform(expectedOutput)
	assert.Equal(test, expectedOutput, InverseTransform(transformation, index))
}

func TestGetIndexes(test *testing.T) {
	assert.Equal(test, []int{1,3,4,2,0}, getIndexes("EPROD", []string{"D", "E", "O","P", "R"}))
}