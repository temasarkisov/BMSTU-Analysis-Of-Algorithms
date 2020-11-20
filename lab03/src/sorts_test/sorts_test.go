package sorts_test

import (
	"testing"

	"../sorts"
)

func BenchmarkBubbleSort(b *testing.B) {
	intSlice := []int{9, 5, 3, 5, 6, 7, 3, 2, 4, 6}
	for i := 0; i < b.N; i++ {
		sorts.BubbleSort(intSlice)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	intSlice := []int{9, 5, 3, 5, 6, 7, 3, 2, 4, 6}
	for i := 0; i < b.N; i++ {
		sorts.InsertionSort(intSlice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	intSlice := []int{9, 5, 3, 5, 6, 7, 3, 2, 4, 6}
	for i := 0; i < b.N; i++ {
		sorts.QuickSort(intSlice)
	}
}
