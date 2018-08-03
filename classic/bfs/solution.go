package main

type Graph interface {
	GetVerticesTotal() int
	AddEdges(vertex1 int, vertex2 int)
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

// TODO implement and test BFS
func BFS(current int, graph Graph) (path []int) {
	visited := make([]bool, graph.GetVerticesTotal())
	queue := make([]int, graph.GetVerticesTotal())

	visited[current] = true
	queue = append([]int{current}, queue...)

	for len(queue) > 0 {
		var item int
		item, queue = queue[0], queue[1:]

		path = append(path, item)
		for _, adj := range graph.GetAdjacencies(item) {
			if !visited[adj] {
				visited[adj] = true
				queue = append([]int{adj}, queue...)
			}
		}
	}

	return
}
