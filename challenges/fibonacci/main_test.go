package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCalculateIterative_BasicCases(t *testing.T) {
	fibo := Fibo{}
	assert.Equal(t, 3, fibo.CalculateIterative(4))
	assert.Equal(t, 0, fibo.CalculateIterative(0))
	assert.Equal(t, 1, fibo.CalculateIterative(1))
	assert.Equal(t, 6765, fibo.CalculateIterative(20))
}

func TestCalculateRecursive_BasicCases(t *testing.T) {
	fibo := Fibo{}
	assert.Equal(t, 3, fibo.CalculateRecursive(4))
	assert.Equal(t, 0, fibo.CalculateRecursive(0))
	assert.Equal(t, 1, fibo.CalculateRecursive(1))
	assert.Equal(t, 6765, fibo.CalculateRecursive(20))
}
