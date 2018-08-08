// This file is not a data structure or algorithm about graph, it's a helper data structure to implement graph algorithms.

package graph

import "container/heap"

// A priorityQueue implements heap.Interface and holds Vertex.
type priorityQueue []*Vertex

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].Dis < pq[j].Dis }

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// 当一个元素的优先级改变时,修正优先队列
func (pq *priorityQueue) fix(item *Vertex) {
	heap.Fix(pq, item.Index)
}
