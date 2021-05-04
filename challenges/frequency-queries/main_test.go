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
