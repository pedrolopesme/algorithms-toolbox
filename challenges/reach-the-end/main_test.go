package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReachTheEnd_Case1(t *testing.T) {
	grid := []string{"..", ".."}
	assert.Equal(t, "Yes", reachTheEnd(grid, 3))
}

func TestReachTheEnd_Case2(t *testing.T) {
	grid := []string{"12", "34"}
	row, column := transformElementToPosition(3, grid)
	assert.Equal(t, int32(1), row)
	assert.Equal(t, int32(1), column)

	row, column = transformElementToPosition(1, grid)
	assert.Equal(t, int32(0), row)
	assert.Equal(t, int32(1), column)

	row, column = transformElementToPosition(0, grid)
	assert.Equal(t, int32(0), row)
	assert.Equal(t, int32(0), column)

	row, column = transformElementToPosition(2, grid)
	assert.Equal(t, int32(1), row)
	assert.Equal(t, int32(0), column)
}

func TestReachTheEnd_Case3(t *testing.T) {
	grid := []string{"12", "34"}
	assert.Equal(t, int32(3), transformPositionToElement(1, 1, grid))
	assert.Equal(t, int32(1), transformPositionToElement(0, 1, grid))
	assert.Equal(t, int32(0), transformPositionToElement(0, 0, grid))
	assert.Equal(t, int32(2), transformPositionToElement(1, 0, grid))
}

func TestReachTheEnd_Case4(t *testing.T) {
	grid := []string{"..##", "#.##", "#..."}
	assert.Equal(t, "Yes", reachTheEnd(grid, int32(5)))
}

func TestReachTheEnd_Case5(t *testing.T) {
	grid := []string{".#", "#."}
	assert.Equal(t, "No", reachTheEnd(grid, int32(5)))
}
