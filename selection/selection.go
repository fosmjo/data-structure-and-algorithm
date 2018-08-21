package selection

import msort "github.com/fosmjo/data-structure-and-algorithm/sort"

// Select 选择第i小的元素
// 使用场景：（1）求中位数，（2）求TopK
// 期望时间复杂度：θ(n)
// 最坏时间复杂度：θ(n^2)
func Select(nums []int, i int) int {
	return selectHelper(nums, 0, len(nums)-1, i)
}

func selectHelper(nums []int, p, r, i int) int {
	if p == r {
		return nums[p]
	}

	q := msort.RandomizedPartition(nums, p, r)
	k := q - p + 1
	if i == k {
		return nums[q]
	} else if i < k {
		return selectHelper(nums, p, q-1, i)
	} else {
		return selectHelper(nums, q+1, r, i-k)
	}
}
