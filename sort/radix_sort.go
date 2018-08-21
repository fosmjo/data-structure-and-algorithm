package sort

type elem struct {
	payload int
	key     int
	digits  []int
}

// RadixSort 基数排序
func RadixSort(nums []int, d int) {
	elems := make([]*elem, len(nums))

	// map
	for i, v := range nums {
		digits := make([]int, d)
		fillDigits(digits, v)
		elems[i] = &elem{payload: v, digits: digits}
	}

	// sort
	for i := 0; i < d; i++ {
		for _, e := range elems {
			e.key = e.digits[i]
		}

		// call counting sort
		elems = countingSort(elems, 10)
	}

	for i, v := range elems {
		nums[i] = v.payload
	}
}

func fillDigits(digits []int, n int) {
	k := 0
	for {
		digit := n % 10
		digits[k] = digit
		k++
		n = n / 10
		if n == 0 {
			break
		}
	}
}

// an ad-hoc counting sort for radix sort
func countingSort(a []*elem, k int) []*elem {
	c := make([]int, k+1)
	for j := 0; j < len(a); j++ {
		c[a[j].key]++
	}
	// c[i] now contains the number of elements equal to i

	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	// c[i] now contains the number of elements less than or equal to i

	b := make([]*elem, len(a))
	for j := len(a) - 1; j >= 0; j-- {
		b[c[a[j].key]-1] = a[j]
		c[a[j].key]--
	}
	return b
}
