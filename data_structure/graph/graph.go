package graph

type Edge struct {
	Src    int
	Dest   int
	Weight float64
}
type Graph struct {
	Vertices int
	Edges    []Edge
}

func InitGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    make([]Edge, 0),
	}
}

func (g *Graph) AddEdge(src, dest int, weight float64) {
	g.Edges = append(g.Edges, Edge{src, dest, weight})
}
