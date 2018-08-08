package graph

import (
	"container/heap"
	"fmt"
)

// Dijkstra 解决带权重的有向图上单源最短路径问题,要求所有边的权重都为非负值
// 时间复杂度:
// 		   1.使用普通二叉堆实现优先队列时为O((|V| + |E|)lg|V|)
//         2.使用斐波那契堆实现优先队列时为O(|E| + |V|lg|V|),适合|V|远小于|E|的情况
func (g *Graph) Dijkstra(vertexName string) error {
	s := g.findVertexByName(vertexName)
	if s == nil {
		return fmt.Errorf("no vertex named %s", vertexName)
	}

	// init
	g.initializeSingleSource(s)
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	for _, v := range g.Vertexes {
		heap.Push(&pq, v)
	}

	// run
	for len(pq) > 0 {
		u := heap.Pop(&pq).(*Vertex)
		for _, v := range g.AdjList[u] {
			g.relax(u, v)
			pq.fix(v)
		}
	}

	return nil
}

func (g *Graph) initializeSingleSource(s *Vertex) {
	for _, v := range g.Vertexes {
		v.Dis = INFINITY
		v.Priv = nil
	}
	s.Dis = 0
}

func (g *Graph) relax(u, v *Vertex) {
	if v.Dis > u.Dis+g.Weights[u][v] {
		v.Dis = u.Dis + g.Weights[u][v]
		v.Priv = u
	}
}
