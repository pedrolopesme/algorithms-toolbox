package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateChangeForZeroValue(test *testing.T) {
	assert.Equal(test, 1, 0)
}

func TestCalculateChangeForExactCoinMatch(test *testing.T) {
	assert.Equal(test, 1, 0)
}

func TestCalculateChangeUsingMultipleCoins(test *testing.T) {
	assert.Equal(test, 1, 0)
}
