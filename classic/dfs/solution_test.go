package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphTraversal(test *testing.T) {
	graph := NewGraph(4)
	graph.AddEdges(0, 1)
	graph.AddEdges(0, 2)
	graph.AddEdges(1, 2)
	graph.AddEdges(2, 0)
	graph.AddEdges(2, 3)
	graph.AddEdges(3, 3)

	assert.Equal(test, []int{2, 0, 1, 3}, DFS(graph, 2))
	assert.Equal(test, []int{0, 1, 2, 3}, DFS(graph, 0))
	assert.Equal(test, []int{3}, DFS(graph, 3))
	assert.Equal(test, []int{1,2,0,3}, DFS(graph, 1))
}
