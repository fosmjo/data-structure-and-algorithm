package graph

import (
	"container/heap"
)

// Prim 用于生成无向图的最小生成树
// 时间复杂度:
// 		   1.使用普通二叉堆实现优先队列时为O(|E|lg|V|)
//         2.使用斐波那契堆实现优先队列时为O(|E| + |V|lg|V|),适合|V|远小于|E|的情况
func (g *Graph) Prim() {
	// init
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	for _, v := range g.Vertexes {
		v.Priv = nil
		v.Dis = INFINITY
		v.InQ = true
		heap.Push(&pq, v)
	}
	g.Vertexes[0].Dis = 0

	// run
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Vertex)
		u.InQ = false

		for _, v := range g.AdjList[u] {
			if v.InQ && g.Weights[u][v] < v.Dis {
				v.Priv = u
				v.Dis = g.Weights[u][v]
				pq.fix(v)
			}
		}
	}
}
