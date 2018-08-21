package sort

// SelectionSort 选择排序
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := nums[i]
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minIndex = j
			}
		}
		if minIndex != i {
			// swap
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}

	}
}
