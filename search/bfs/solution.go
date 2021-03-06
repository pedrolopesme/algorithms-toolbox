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
		edges: make(map[int][]int),
	}
}

func BFS(current int, graph Graph) (path []int) {
	queue := []int{current}
	visited := make([]bool, graph.GetVerticesTotal())
	visited[current] = true

	for len(queue) > 0 {
		var item int
		item, queue = queue[0], queue[1:]
		path = append(path, item)
		for _, adj := range graph.GetAdjacencies(item) {
			if !visited[adj] {
				visited[adj] = true
				queue = append(queue, adj)
			}
		}
	}

	return
}
