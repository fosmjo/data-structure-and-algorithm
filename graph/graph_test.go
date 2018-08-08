package graph

import (
	"fmt"
	"strconv"
	"testing"
)

// 测试数据数据见<<数据结构与算法分析>>相关章节

// 有环
func testData() [][3]string {
	edges := [][3]string{
		[3]string{"v1", "v2", "2"},
		[3]string{"v1", "v4", "1"},
		[3]string{"v2", "v4", "3"},
		[3]string{"v2", "v5", "10"},
		[3]string{"v3", "v1", "4"},
		[3]string{"v3", "v6", "5"},
		[3]string{"v4", "v3", "2"},
		[3]string{"v4", "v5", "2"},
		[3]string{"v4", "v6", "8"},
		[3]string{"v4", "v7", "4"},
		[3]string{"v5", "v7", "6"},
		[3]string{"v7", "v6", "1"},
	}
	return edges
}

// 无环
func testData2() [][3]string {
	edges := [][3]string{
		[3]string{"v1", "v2", "2"},
		[3]string{"v1", "v3", "4"},
		[3]string{"v1", "v4", "1"},
		[3]string{"v2", "v4", "3"},
		[3]string{"v2", "v5", "10"},
		[3]string{"v3", "v6", "5"},
		[3]string{"v4", "v3", "2"},
		[3]string{"v4", "v6", "8"},
		[3]string{"v4", "v7", "4"},
		[3]string{"v5", "v4", "7"},
		[3]string{"v5", "v7", "6"},
		[3]string{"v7", "v6", "1"},
	}
	return edges
}

func newGraph(edges [][3]string, isDirected bool) *Graph {
	g := NewGraph(isDirected)
	for _, edge := range edges {
		w, _ := strconv.Atoi(edge[2])
		g.AddEdge(edge[0], edge[1], w)
	}
	return g
}

func TestBFS(t *testing.T) {
	edges := testData()
	g := newGraph(edges, true)
	err := g.BFS("v3")
	fmt.Println("BFS result:")
	if err == nil {
		for _, v := range g.Vertexes {
			fmt.Printf("v3 -> %s : %d\n", v.Name, v.Dis)
		}
	} else {
		t.Error(err)
	}
}

func TestDFS(t *testing.T) {
	edges := testData()
	g := newGraph(edges, true)
	g.DFS()

	fmt.Println("DFS result:")
	for _, v := range g.Vertexes {
		fmt.Printf("%s : %d\n", v.Name, v.Dis)
	}
}

func TestTopologicalSort(t *testing.T) {
	edges := testData2()
	g := newGraph(edges, true)
	result, err := g.TopologicalSort()
	fmt.Println("TopologicalSort result:")
	if err == nil {
		for _, v := range result {
			fmt.Print(v.Name, " ")
		}
		fmt.Println()
	} else {
		t.Error(err)
	}
}

func TestTopologicalSort2(t *testing.T) {
	edges := testData2()
	g := newGraph(edges, true)
	result := g.TopologicalSort2()
	fmt.Println("TopologicalSort2 result:")
	for _, v := range result {
		fmt.Print(v.Name, " ")
	}
	fmt.Println()
}

func TestPrim(t *testing.T) {
	edges := testData2()
	g := newGraph(edges, false)
	g.Prim()
	fmt.Println("Prim result:")
	for i, v := range g.Vertexes {
		if i != 0 {
			fmt.Println(v.Priv.Name, " -- ", v.Name)
		}
	}
}

func TestKruskal(t *testing.T) {
	edges := testData2()
	g := newGraph(edges, false)
	result := g.Kruskal()
	fmt.Println("Kruskal result:")
	for _, e := range result {
		fmt.Println(e.Start.Name, " -- ", e.End.Name)
	}
}

func TestDijkstra(t *testing.T) {
	edges := testData()
	g := newGraph(edges, true)
	err := g.Dijkstra("v1")
	fmt.Println("Dijkstra result:")
	if err == nil {
		for _, v := range g.Vertexes {
			fmt.Println("v1", " -> ", v.Name, " : ", v.Dis)
		}
	} else {
		t.Error(err)
	}
}

func TestBellmanFord(t *testing.T) {
	edges := testData()
	g := newGraph(edges, true)
	exist, err := g.BellmanFord("v1")
	fmt.Println("BellmanFord result:")
	if exist {
		for _, v := range g.Vertexes {
			fmt.Println("v1", " -> ", v.Name, " : ", v.Dis)
		}
	} else {
		t.Error(err)
	}
}
