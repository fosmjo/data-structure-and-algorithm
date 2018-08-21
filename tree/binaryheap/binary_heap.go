package binaryheap

import (
	"errors"
)

// BinaryHeap 二叉堆，是一棵完全二叉树
type BinaryHeap struct {
	Data []int
	Size int
	Cap  int
}

// IsFull 满了
func (bh *BinaryHeap) IsFull() bool {
	return bh.Size == bh.Cap
}

// IsEmpty 空了
func (bh *BinaryHeap) IsEmpty() bool {
	return bh.Size == 0
}

// Insert 插入
func (bh *BinaryHeap) Insert(x int) error {
	if bh.IsFull() {
		return errors.New("heap is full")
	}

	bh.Size++
	bh.Data[bh.Size] = x
	bh.percolateUp(bh.Size)

	return nil
}

// percolateUp 上渗
func (bh *BinaryHeap) percolateUp(i int) {
	x := bh.Data[i]
	for ; bh.Data[i/2] > x && i > 1; i /= 2 {
		bh.Data[i] = bh.Data[i/2]
	}
	bh.Data[i] = x
}

// DeleteMin 删除最小
func (bh *BinaryHeap) DeleteMin() (int, error) {
	if bh.IsEmpty() {
		return -1, errors.New("heap is empty")
	}

	min := bh.Data[1]
	bh.Data[1] = bh.Data[bh.Size]
	bh.Data[bh.Size] = 0
	bh.Size--
	bh.percolateDown(1)
	return min, nil
}

// percolateDown 下滤, 维护以i为根节点的子树以满足最小堆性质（min_heapify）
func (bh *BinaryHeap) percolateDown(i int) {
	x := bh.Data[i]
	for {
		left, right := 2*i, 2*i+1
		if bh.Size < left {
			bh.Data[i] = x
			break
		} else if bh.Size == left {
			if x > bh.Data[left] {
				// swap
				bh.Data[i], bh.Data[left] = bh.Data[left], x
			} else {
				bh.Data[i] = x
			}
			break
		} else {
			if x <= bh.Data[left] && x <= bh.Data[right] {
				bh.Data[i] = x
				break
			} else if bh.Data[left] < bh.Data[right] {
				bh.Data[i], i = bh.Data[left], left
			} else {
				bh.Data[i], i = bh.Data[right], right
			}
		}
	}
}

// percolateDownRecursive 下滤,递归版本
func (bh *BinaryHeap) percolateDownRecursive(i int) {
	minIndex, left, right := i, 2*i, 2*i+1

	if left <= bh.Size && bh.Data[left] < bh.Data[i] {
		minIndex = left
	}
	if right <= bh.Size && bh.Data[right] < bh.Data[minIndex] {
		minIndex = right
	}

	if minIndex != i {
		// swap
		bh.Data[minIndex], bh.Data[i] = bh.Data[i], bh.Data[minIndex]
		// advance
		bh.percolateDownRecursive(minIndex)
	}
}

// NewBinaryHeap 构造函数
func NewBinaryHeap(nums []int, cap int) *BinaryHeap {
	data := make([]int, cap+1)
	bh := &BinaryHeap{Data: data, Size: 0, Cap: cap}
	// build heap
	for _, n := range nums {
		bh.Insert(n)
	}
	return bh
}
