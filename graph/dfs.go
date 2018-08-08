package graph

// 发现节点或者访问完节点的时间
var time int

// DFS 广度优先搜索
// 时间复杂度: O(|V| + |E|)
func (g *Graph) DFS() {
	// init
	for _, u := range g.Vertexes {
		u.Color = WHITE
		u.Priv = nil
	}

	time = 0
	for _, u := range g.Vertexes {
		if u.Color == WHITE {
			g.dfsVisit(u)
		}
	}
}

func (g *Graph) dfsVisit(u *Vertex) {
	time++
	u.Dis = time
	u.Color = GRAY

	for _, v := range g.AdjList[u] {
		if v.Color == WHITE {
			v.Priv = u
			g.dfsVisit(v)
		}
	}

	u.Color = BLACK
	time++
	u.Fin = time
}
