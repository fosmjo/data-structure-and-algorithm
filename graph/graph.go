package graph

const (
	INFINITY = 10000

	WHITE Color = iota
	GRAY
	BLACK
)

// Color 顶点颜色
type Color int

// Vertex 顶点
type Vertex struct {
	Name  string
	Dis   int
	Fin   int
	Priv  *Vertex
	Color Color
	Index int
	InQ   bool
}

// Edge 边
type Edge struct {
	Start  *Vertex
	End    *Vertex
	Weight int
}

// Graph 邻接表表示的图
type Graph struct {
	Vertexes   []*Vertex
	Edges      []*Edge
	AdjList    map[*Vertex][]*Vertex
	Weights    map[*Vertex]map[*Vertex]int
	IsDirected bool
}

// NewGraph 构造函数
func NewGraph(isDirected bool) *Graph {
	return &Graph{
		Vertexes:   make([]*Vertex, 0),
		Edges:      make([]*Edge, 0),
		AdjList:    make(map[*Vertex][]*Vertex),
		Weights:    make(map[*Vertex]map[*Vertex]int),
		IsDirected: isDirected,
	}
}

// AddEdge 添加边
func (g *Graph) AddEdge(start string, end string, weight int) {
	s := g.findOrCreateVertexByName(start)
	e := g.findOrCreateVertexByName(end)
	g.addEdge(s, e, weight)
	if !g.IsDirected {
		g.addEdge(e, s, weight)
	}
}

func (g *Graph) addEdge(start *Vertex, end *Vertex, weight int) {
	g.AdjList[start] = append(g.AdjList[start], end)
	g.Edges = append(g.Edges, &Edge{Start: start, End: end, Weight: weight})

	if g.Weights[start] == nil {
		g.Weights[start] = make(map[*Vertex]int)
	}
	g.Weights[start][end] = weight
}

func (g *Graph) findOrCreateVertexByName(name string) *Vertex {
	v := g.findVertexByName(name)

	if v == nil {
		v = &Vertex{
			Name:  name,
			Color: WHITE,
		}
		g.Vertexes = append(g.Vertexes, v)
	}

	return v
}

func (g *Graph) findVertexByName(name string) *Vertex {
	for _, v := range g.Vertexes {
		if v.Name == name {
			return v
		}
	}
	return nil
}
