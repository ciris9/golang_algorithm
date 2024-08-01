package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	// 创建一个示例图
	g := Graph{
		Vertices: 8,
		Edges: []Edge{
			{0, 1, 4}, {0, 7, 8}, {1, 2, 8}, {1, 7, 11}, {1, 6, 4},
			{2, 3, 2}, {2, 5, 4}, {2, 6, 6}, {2, 7, 7}, {3, 4, 6},
			{3, 5, 14}, {4, 5, 9}, {4, 6, 2}, {5, 6, 10}, {5, 7, 6},
			{6, 7, 1},
		},
	}

	// 执行 Dijkstra 算法
	dist := Dijkstra(&g, 0)

	fmt.Println("Shortest distances from source vertex 0:")
	for _, v := range dist {
		fmt.Printf("Vertex %d: %d\n", v.ID, v.Dist)
	}
}
