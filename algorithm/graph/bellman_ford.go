package graph

import (
	"math"
)

// BellmanFord 执行 Bellman-Ford 算法计算单源最短路径
func BellmanFord(g *Graph, source int) ([]int, bool) {
	// 初始化距离数组，将所有距离设置为正无穷
	dist := make([]int, g.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0 // 源点到自身的距离为 0

	// 进行 V-1 次迭代，更新每个顶点的最短距离
	for i := 0; i < g.Vertices-1; i++ {
		for _, e := range g.Edges {
			if dist[e.Source]+e.Weight < dist[e.Dest] {
				dist[e.Dest] = dist[e.Source] + e.Weight
			}
		}
	}

	// 检查是否存在负权重环
	for _, e := range g.Edges {
		if dist[e.Source]+e.Weight < dist[e.Dest] {
			return nil, false // 存在负权重环
		}
	}

	return dist, true
}
