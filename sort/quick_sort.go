package sort

import "math/rand"

// QuickSort 快速排序
func QuickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, p, r int) {
	if p < r {
		q := RandomizedPartition(nums, p, r)
		quickSortHelper(nums, p, q-1)
		quickSortHelper(nums, q+1, r)
	}
}

// RandomizedPartition 随机划分
func RandomizedPartition(nums []int, p, r int) int {
	i := p + rand.Intn(r-p+1)
	nums[i], nums[r] = nums[r], nums[i] // swap
	return partition(nums, p, r)
}

func partition(nums []int, p, r int) int {
	pivot, i := nums[r], p-1

	for j := p; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i] // swap
		}
	}
	// for any index ∈ [p, r]
	// if index <=i, then num[index] <= pivot
	// if index > i, then num[index] >  pivot
	// here we make the pivot to split nums
	i++
	nums[i], nums[r] = nums[r], nums[i] // swap
	return i
}
