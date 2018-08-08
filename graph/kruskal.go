package graph

import (
	"sort"
)

// Kruskal 用于生成无向图的最小生成树
// 时间复杂度: O(|E|lg|V|)
func (g *Graph) Kruskal() []*Edge {
	edges := make([]*Edge, 0)
	set := &disjointSet{make(map[*Vertex]*Vertex)}
	for _, v := range g.Vertexes {
		set.makeSet(v)
	}

	sort.Slice(g.Edges, func(i, j int) bool { return g.Edges[i].Weight < g.Edges[j].Weight })

	for _, e := range g.Edges {
		u, v := e.Start, e.End
		if set.findSet(u) != set.findSet(v) {
			edges = append(edges, e)
			set.union(u, v)
		}
	}

	return edges
}
