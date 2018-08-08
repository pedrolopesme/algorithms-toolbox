package main

type Graph interface {
	GetVerticesTotal() int
	AddEdges(vertex1 int, vertex2 int)
	GetEdges() map[int][]int
	GetAdjacencies(vertex int) []int
}

type BaseGraph struct {
	// Total of vertices
	vertices int

	// Edges between vertices, eg. vertex1 -> [vertex2, vertex3, ...]
	edges map[int][]int
}

func (g BaseGraph) GetVerticesTotal() int {
	return g.vertices
}

func (g BaseGraph) GetEdges() map[int][]int {
	return g.edges
}

func (g BaseGraph) AddEdges(vertex1 int, vertex2 int) {
	g.edges[vertex1] = append(g.edges[vertex1], vertex2)
}

func (g BaseGraph) GetAdjacencies(vertex int) []int {
	return g.edges[vertex]
}

func NewGraph(totalOfVertices int) *BaseGraph {
	return &BaseGraph{
		vertices: totalOfVertices,
		edges:    make(map[int][]int),
	}
}

func DFS(G Graph, vertex int, visited *map[int]bool, path *[]int) {
	(*visited)[vertex] = true
	*path = append(*path, vertex)
	for _, adj := range G.GetAdjacencies(vertex) {
		if !(*visited)[adj] {
			DFS(G, adj, visited, path)
		}
	}
}

// RunDFS just make it easier to run DFS
func RunDFS(G Graph, vertex int) (path []int) {
	visited := make(map[int]bool, G.GetVerticesTotal())
	DFS(G, vertex, &visited, &path)
	return
}
