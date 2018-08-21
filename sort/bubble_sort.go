package sort

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
