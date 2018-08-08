package graph

import (
	"errors"
	"fmt"
)

// BellmanFord 解决一般情况下的单源最短路径问题,边的权重可以为负值,
// 如果存在权重为负值的环路,算法将告诉我们不存在解决方案
// 时间复杂度: O(|V||E|)
func (g *Graph) BellmanFord(vertexName string) (bool, error) {
	s := g.findVertexByName(vertexName)
	if s == nil {
		return false, fmt.Errorf("no vertex named %s", vertexName)
	}

	g.initializeSingleSource(s)

	for i := 1; i < len(g.Vertexes); i++ {
		for _, e := range g.Edges {
			g.relax(e.Start, e.End)
		}
	}

	for _, e := range g.Edges {
		u, v := e.Start, e.End

		if v.Dis > u.Dis+g.Weights[u][v] {
			return false, errors.New("there is a loop with negtive weight")
		}
	}
	return true, nil
}
