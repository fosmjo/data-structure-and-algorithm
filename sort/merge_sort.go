package sort

// MergeSort 归并排序
func MergeSort(nums []int) {
	mergeSortHelper(nums, 0, len(nums)-1)
}

func mergeSortHelper(nums []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSortHelper(nums, p, q)
		mergeSortHelper(nums, q+1, r)
		merge(nums, p, q, r)
	}
}

func merge(nums []int, p, q, r int) {
	left := make([]int, q-p+1)
	copy(left, nums[p:q+1])

	right := make([]int, r-q)
	copy(right, nums[q+1:r+1])

	k, i, j := p, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}

	for ; i < len(left); i++ {
		nums[k] = left[i]
		k++
	}

	for ; j < len(right); j++ {
		nums[k] = right[j]
		k++
	}
}
