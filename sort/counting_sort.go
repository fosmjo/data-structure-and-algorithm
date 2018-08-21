package sort

// CountingSort 计数排序
// 适用场景：输入元素(或者元素的Key)都属于[0,k]
func CountingSort(a []int, k int) []int {
	c := make([]int, k+1)
	for j := 0; j < len(a); j++ {
		c[a[j]]++
	}
	// c[i] now contains the number of elements equal to i

	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	// c[i] now contains the number of elements less than or equal to i

	b := make([]int, len(a))
	for j := len(a) - 1; j >= 0; j-- {
		b[c[a[j]]-1] = a[j]
		c[a[j]]--
	}
	return b
}
