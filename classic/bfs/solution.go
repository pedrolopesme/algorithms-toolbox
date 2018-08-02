package main

type Graph interface {
	AddEdges(vertex1 int, vertex2 int)
	GetAdjacencies(vertex int) []int
}

type BaseGraph struct {
	nodes map[int][]int
}

func (g BaseGraph) AddEdges(vertex1 int, vertex2 int) {
	g.nodes[vertex1] = append(g.nodes[vertex1], vertex2)
}

func (g BaseGraph) GetAdjacencies(vertex int) []int {
	return g.nodes[vertex]
}

func NewGraph() *BaseGraph {
	return &BaseGraph{
		nodes: make(map[int][]int),
	}
}