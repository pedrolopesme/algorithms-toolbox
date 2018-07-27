package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
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
	input := "TEXYDST.E.IXIXIXXSSMPPS.B..E.S.EUSFXDIIOIIIT"
	expectedOutput := "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES"
	index, tranformation := Transform(input)

	fmt.Println(index)

	assert.Equal(test, expectedOutput, InverseTransform(tranformation, index))
}