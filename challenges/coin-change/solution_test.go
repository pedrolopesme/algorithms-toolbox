package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateChangeForZeroValue(test *testing.T) {
	coins := map[int]int{50: 100, 25: 100, 10: 100, 5: 100, 1: 100}
	result := Solution(coins, 0.0)
	assert.Equal(test, len(result), 0)
}

func TestCalculateChangeForExactCoinMatch(test *testing.T) {
	coins := map[int]int{50: 100, 25: 100, 10: 100, 5: 100, 1: 100}
	result := Solution(coins, 0.01)
	assert.Equal(test, len(result), 1)
	assert.Equal(test, result[1], 1)
}

// TODO fix broken test
func TestCalculateChangeUsingMultipleCoins(test *testing.T) {
	coins := map[int]int{50: 100, 25: 100, 10: 100, 5: 100, 1: 100}
	result := Solution(coins, 1.17)
	assert.Equal(test, len(result), 6)
	assert.Equal(test, result[50], 2)
	assert.Equal(test, result[10], 1)
	assert.Equal(test, result[5], 1)
	assert.Equal(test, result[1], 2)
}
