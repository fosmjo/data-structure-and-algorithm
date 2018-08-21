package sort

import (
	"reflect"
	"testing"
)

func testData() ([]int, []int) {
	nums := []int{5, 2, 4, 6, 1, 3}
	result := []int{1, 2, 3, 4, 5, 6}
	return nums, result
}

func TestSelectionSort(t *testing.T) {
	nums, result := testData()
	SelectionSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("SelectionSort failed")
	}
}

func TestBubbleSort(t *testing.T) {
	nums, result := testData()
	BubbleSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("BubbleSort failed")
	}
}

func TestInsertionSort(t *testing.T) {
	nums, result := testData()
	InsertionSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("InsertionSort failed")
	}
}

func TestMergeSort(t *testing.T) {
	nums, result := testData()
	MergeSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("MergeSort failed")
	}
}

func TestHeapSort(t *testing.T) {
	nums, result := testData()
	HeapSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("HeapSort failed")
	}
}
func TestQuickSort(t *testing.T) {
	nums, result := testData()
	QuickSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Error("QuickSort failed")
	}
}

func TestCountingSort(t *testing.T) {
	nums := []int{2, 5, 3, 0, 2, 3, 0, 3}
	expectResult := []int{0, 0, 2, 2, 3, 3, 3, 5}
	realResult := CountingSort(nums, 5)
	if !reflect.DeepEqual(realResult, expectResult) {
		t.Error("CountingSort failed")
	}
}

func TestRadixSort(t *testing.T) {
	nums := []int{329, 457, 657, 839, 436, 720, 355}
	expectResult := []int{329, 355, 436, 457, 657, 720, 839}
	RadixSort(nums, 3)
	if !reflect.DeepEqual(nums, expectResult) {
		t.Error("RadixSort failed")
	}
}

func TestBucketSort(t *testing.T) {
	nums := []float64{0.79, 0.13, 0.16, 0.64, 0.39, 0.2, 0.89, 0.53, 0.71, 0.42}
	expectResult := []float64{0.13, 0.16, 0.2, 0.39, 0.42, 0.53, 0.64, 0.71, 0.79, 0.89}
	BucketSort(nums)
	if !reflect.DeepEqual(nums, expectResult) {
		t.Error("BucketSort failed")
	}
}
