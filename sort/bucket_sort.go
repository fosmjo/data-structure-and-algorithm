package sort

// BucketSort 桶排序
// 适用场景：元素均匀、独立地分布在区间[0,1)上
func BucketSort(nums []float64) {
	n := len(nums)
	buckets := make([][]float64, n)
	for i := range buckets {
		buckets[i] = make([]float64, 0)
	}

	for i := range nums {
		index := int(float64(n) * nums[i])
		buckets[index] = append(buckets[index], nums[i])
	}

	for i := range buckets {
		insertionSort(buckets[i])
	}

	k := 0
	for _, b := range buckets {
		for _, v := range b {
			nums[k] = v
			k++
		}
	}
}

// an ad-hoc insertion sort for bucket sort
func insertionSort(nums []float64) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > key; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = key
	}
}
