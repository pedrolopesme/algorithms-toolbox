package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreqQuery_Case1(t *testing.T) {
	queries := [][]int32{}
	queries = append(queries, []int32{1, 1})
	queries = append(queries, []int32{2, 2})
	queries = append(queries, []int32{3, 2})
	queries = append(queries, []int32{1, 1})
	queries = append(queries, []int32{1, 1})
	queries = append(queries, []int32{2, 1})
	queries = append(queries, []int32{3, 2})
	assert.Equal(t, []int32{0, 1}, freqQuery(queries))
}

func TestFreqQuery_Case2(t *testing.T) {
	queries := [][]int32{}
	queries = append(queries, []int32{1, 5})
	queries = append(queries, []int32{1, 6})
	queries = append(queries, []int32{3, 2})
	queries = append(queries, []int32{1, 10})
	queries = append(queries, []int32{1, 10})
	queries = append(queries, []int32{1, 6})
	queries = append(queries, []int32{2, 5})
	queries = append(queries, []int32{3, 2})
	assert.Equal(t, []int32{0, 1}, freqQuery(queries))
}

func TestFreqQuery_Case3(t *testing.T) {
	queries := [][]int32{}
	queries = append(queries, []int32{3, 4})
	queries = append(queries, []int32{2, 1003})
	queries = append(queries, []int32{1, 16})
	queries = append(queries, []int32{3, 1})
	assert.Equal(t, []int32{0, 1}, freqQuery(queries))
}

func TestFreqQuery_Case4(t *testing.T) {
	queries := [][]int32{}
	queries = append(queries, []int32{1, 3})
	queries = append(queries, []int32{2, 3})
	queries = append(queries, []int32{3, 2})
	queries = append(queries, []int32{1, 4})
	queries = append(queries, []int32{1, 5})
	queries = append(queries, []int32{1, 5})
	queries = append(queries, []int32{1, 4})
	queries = append(queries, []int32{3, 2})
	queries = append(queries, []int32{2, 4})
	queries = append(queries, []int32{3, 2})
	assert.Equal(t, []int32{0, 1, 1}, freqQuery(queries))
}
