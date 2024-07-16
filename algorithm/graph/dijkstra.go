package graph

import "math"

var edge [2505][2505]int
var dis [2505]int
var book [2505]int
var n, s int
var inf = math.MaxInt/2 - 1

func Dijkstra() {
	for i := 1; i <= n; i++ {
		dis[i] = edge[s][i]
	}
	book[s] = 1
	for i := 1; i <= n-1; i++ {
		var minn int = inf
		var nextPos int = -1
		for j := 1; j <= n; j++ {
			if book[j] == 0 && dis[j] < minn {
				nextPos = j
				minn = dis[j]
			}
		}
		if nextPos == -1 {
			break
		}
		book[nextPos] = 1
		for j := 1; j <= n; j++ {
			if book[j] == 0 && dis[nextPos]+edge[nextPos][j] < dis[j] {
				dis[j] = dis[nextPos] + edge[nextPos][j]
			}
		}
	}
}
