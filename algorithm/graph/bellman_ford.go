package graph

import (
	"algorithm/data_structure/graph"
	"fmt"
	"math"
)

func BellmanFord(g *graph.Graph, source int) ([]float64, []int) {
	dist := make([]float64, g.Vertices)
	prev := make([]int, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		dist[i] = math.Inf(1) // 将所有节点的距离设为 无穷大
		prev[i] = -1          // 保存到节点v最短距离，到v的上一跳是哪个节点
	}
	dist[source] = 0 // 起点的距离为 0

	// Relax Edges V - 1 times（i 从 1 开始计数）
	for i := 1; i < g.Vertices; i++ {
		for _, edge := range g.Edges {
			u := edge.Src
			v := edge.Dest
			w := edge.Weight
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
			}
		}
	}

	// Check for negative cycle，经过了 v-1 次循环了，还是存在更短的路径证明存在负环了
	for _, edge := range g.Edges {
		u := edge.Src
		v := edge.Dest
		w := edge.Weight
		if dist[u]+w < dist[v] {
			fmt.Println("Graph contains a negative weight cycle")
			return nil, nil
		}
	}

	return dist, prev
}

func PrintShortestPaths(dist []float64, prev []int, source int) {
	fmt.Println("Shortest Paths from vertex", source)
	for i := 0; i < len(dist); i++ {
		if dist[i] == math.Inf(1) {
			fmt.Printf("Vertex %d is not reachable\n", i)
		} else {
			var path []int
			j := i
			for j != -1 {
				path = append([]int{j}, path...)
				j = prev[j]
			}
			fmt.Printf("Vertex %d: Distance=%f, Path=%v\n", i, dist[i], path)
		}
	}
}
