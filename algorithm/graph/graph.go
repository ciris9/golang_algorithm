package graph

// Edge 表示图中的一条边
type Edge struct {
	Source int
	Dest   int
	Weight int
}

// Graph 表示图结构
type Graph struct {
	Vertices int
	Edges    []Edge
}
