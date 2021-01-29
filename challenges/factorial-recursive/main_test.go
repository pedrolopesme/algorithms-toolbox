package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFactorialRecursive_BasicCases(t *testing.T) {
	f := FactorialRecursive{}
	assert.Equal(t, 24, f.Calculate(4))
	assert.Equal(t, 362880, f.Calculate(9))
	assert.Equal(t, 1, f.Calculate(0))
}
