package binaryheap

import (
	"testing"
)

func newBinaryHeap() *BinaryHeap {
	data := []int{13, 21, 16, 24, 31, 19, 68, 65, 26, 32}
	bh := NewBinaryHeap(data, 20)
	return bh
}

// 检查二叉堆是否符合最小堆性质
func checkBinaryHeap(t *testing.T, bh *BinaryHeap) {
	for i := 1; i <= bh.Size/2; i++ {
		left, right := i*2, i*2+1
		if left <= bh.Size && !(bh.Data[i] <= bh.Data[left]) {
			t.Errorf("left child %d is bigger than its parent, heap: %#v", left, bh)
			break
		}
		if right <= bh.Size && !(bh.Data[i] <= bh.Data[right]) {
			t.Errorf("right child %d is bigger than its parent, heap: %#v", right, bh)
		}
	}
}

func TestBinaryHeap(t *testing.T) {
	bh := newBinaryHeap()
	checkBinaryHeap(t, bh)
}

func TestDeleteMin(t *testing.T) {
	bh := newBinaryHeap()
	for bh.Size > 0 {
		_, err := bh.DeleteMin()
		if err != nil {
			t.Errorf("err: %s, heap:%#v", err, bh)
		}
		checkBinaryHeap(t, bh)
	}
}
