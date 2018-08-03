package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSStartingFromVertexZero(test *testing.T) {
	graph := NewGraph(4)
	graph.AddEdges(0, 1)
	graph.AddEdges(0, 2)
	graph.AddEdges(1, 2)
	graph.AddEdges(2, 0)
	graph.AddEdges(2, 3)
	graph.AddEdges(3, 3)

	path := BFS(2, graph)
	assert.Equal(test,[]int{2, 0, 3, 1}, path)
}

