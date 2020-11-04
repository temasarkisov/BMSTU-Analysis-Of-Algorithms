package sorts

import "math/rand"

// BubbleSort sorts a slice of integers in increasing order with bubble sort algorithm
func BubbleSort(intSlice []int) {
	for i := len(intSlice); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
	}
}

func swap(intSlice []int, i int) {
	temp := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = temp
}

// InsertionSort sorts a slice of integers in increasing order with insertion sort algorithm
func InsertionSort(intSlice []int) {
	var n = len(intSlice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if intSlice[j-1] > intSlice[j] {
				intSlice[j-1], intSlice[j] = intSlice[j], intSlice[j-1]
			}
			j = j - 1
		}
	}
}

// QuickSort sorts a slice of integers in increasing order with quick sort algorithm
func QuickSort(intSlice []int) []int {
	if len(intSlice) < 2 {
		return intSlice
	}
	left, right := 0, len(intSlice)-1
	pivot := rand.Int() % len(intSlice)
	intSlice[pivot], intSlice[right] = intSlice[right], intSlice[pivot]
	for i := range intSlice {
		if intSlice[i] < intSlice[right] {
			intSlice[left], intSlice[i] = intSlice[i], intSlice[left]
			left++
		}
	}
	intSlice[left], intSlice[right] = intSlice[right], intSlice[left]
	QuickSort(intSlice[:left])
	QuickSort(intSlice[left+1:])
	return intSlice
}
