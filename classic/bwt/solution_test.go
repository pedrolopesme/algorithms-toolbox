package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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