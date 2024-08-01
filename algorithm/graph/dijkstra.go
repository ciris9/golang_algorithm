package graph

import (
	"container/heap"
	"math"
)

// PriorityQueue 用于实现 Dijkstra 算法
type PriorityQueue []*Vertex

type Vertex struct {
	ID    int
	Dist  int
	Prev  *Vertex
	Index int // 用于在 PriorityQueue 中的索引
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	v := x.(*Vertex)
	v.Index = n
	*pq = append(*pq, v)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	v := old[n-1]
	v.Index = -1 // for safety
	*pq = old[0 : n-1]
	return v
}

func Dijkstra(g *Graph, source int) []*Vertex {
	dist := make([]*Vertex, g.Vertices)
	for i := range dist {
		dist[i] = &Vertex{ID: i, Dist: math.MaxInt32, Prev: nil}
	}
	dist[source] = &Vertex{ID: source, Dist: 0, Prev: nil}
	pq := make(PriorityQueue, 0, g.Vertices)
	heap.Init(&pq)
	heap.Push(&pq, dist[source])

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Vertex)
		if u.Dist == math.MaxInt32 {
			break
		}

		for _, e := range g.Edges {
			if e.Source == u.ID {
				v := dist[e.Dest]
				if v.Dist > u.Dist+e.Weight {
					v.Dist = u.Dist + e.Weight
					v.Prev = u
					if v.Index != -1 {
						heap.Fix(&pq, v.Index)
					} else {
						heap.Push(&pq, v)
					}
				}
			}
		}
	}

	return dist
}
