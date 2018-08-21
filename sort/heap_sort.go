package sort

import "github.com/fosmjo/data-structure-and-algorithm/tree/binaryheap"

// HeapSort 堆排序
func HeapSort(nums []int) {
	bh := binaryheap.NewBinaryHeap(nums, len(nums))
	k := 0
	for i := bh.Size; i > 0; i-- {
		min, _ := bh.DeleteMin()
		nums[k] = min
		k++
	}
}
