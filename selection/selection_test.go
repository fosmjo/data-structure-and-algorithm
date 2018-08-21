package selection

import "testing"

func TestSelect(t *testing.T) {
	nums := []int{5, 2, 4, 6, 1, 3}
	val := Select(nums, 3)
	if val != 3 {
		t.Errorf("failed to find 3rd element")
	}
}
