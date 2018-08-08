package graph

import (
	"errors"
	"sort"
)

// TopologicalSort 拓扑排序, 适用于有向无环图(DAG)
// 时间复杂度: O(|V| + |E|)
// Kahn's algorithm
func (g *Graph) TopologicalSort() ([]*Vertex, error) {
	result := make([]*Vertex, 0, len(g.Vertexes))
	indegrees := g.indegrees()
	zeroIndegreeVertexes := make([]*Vertex, 0)

	for v, indegree := range indegrees {
		if indegree == 0 {
			zeroIndegreeVertexes = append(zeroIndegreeVertexes, v)
		}
	}

	for len(zeroIndegreeVertexes) > 0 {
		// pop
		u := zeroIndegreeVertexes[0]
		zeroIndegreeVertexes = zeroIndegreeVertexes[1:]

		result = append(result, u)

		for _, v := range g.AdjList[u] {
			indegrees[v]--
			if indegrees[v] == 0 {
				zeroIndegreeVertexes = append(zeroIndegreeVertexes, v)
			}
		}
	}

	if len(result) != cap(result) {
		return nil, errors.New("graph has a cycle")
	}

	return result, nil
}

// 计算每个顶点的入度
func (g *Graph) indegrees() map[*Vertex]int {
	indegrees := make(map[*Vertex]int, len(g.Vertexes))

	for k, vs := range g.AdjList {
		if _, ok := indegrees[k]; !ok {
			indegrees[k] = 0
		}

		for _, v := range vs {
			indegrees[v]++
		}
	}

	return indegrees
}

// TopologicalSort2 拓扑排序, 适用于有向无环图(DAG)
// 时间复杂度: O(|V| + |E|)
// based on DFS
func (g *Graph) TopologicalSort2() []*Vertex {
	g.DFS()
	topoSeq := make([]*Vertex, len(g.Vertexes))
	copy(topoSeq, g.Vertexes)
	sort.Slice(topoSeq, func(i, j int) bool { return topoSeq[i].Fin > topoSeq[j].Fin })
	return topoSeq
}
