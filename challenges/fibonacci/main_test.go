package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCalculateIterative_BasicCases(t *testing.T) {
	fibo := Fibo{}
	assert.Equal(t, 3, fibo.CalculateIterative(4))
}
