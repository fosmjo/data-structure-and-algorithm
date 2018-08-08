package graph

import "fmt"

// BFS 广度优先搜索
// 时间复杂度: O(|V| + |E|)
func (g *Graph) BFS(vertexName string) error {
	s := g.findOrCreateVertexByName(vertexName)
	if s == nil {
		return fmt.Errorf("no vertex named %s", vertexName)
	}

	s.Color = GRAY
	queue := []*Vertex{s}

	for len(queue) > 0 {
		// pop
		u := queue[0]
		queue = queue[1:]

		for _, v := range g.AdjList[u] {
			if v.Color == WHITE {
				v.Color = GRAY // found
				v.Dis = u.Dis + 1
				v.Priv = u
				queue = append(queue, v)
			}
		}
		u.Color = BLACK // visited
	}

	return nil
}
