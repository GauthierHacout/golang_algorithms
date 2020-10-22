package algorithms

import "math/rand"

// Because of the nature of slice in Golang this implementation will actually update the []int passed as a parameter
//as well as returning the sorted slice
func QuickSort(A []int) {
	if len(A) < 2 {
		return
	}

	array := A
	// Bounds of the array (index)
	left, right := 0, len(array)-1
	// Create a pivot
	pivot := rand.Int()% len(array)
	// Swap values for pivot and right index
	array[pivot], array[right] = array[right], array[pivot]
	// Swap values of left until every values which index are smaller than pivot are smaller than array[pivot]
	for i, _ := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	// Swap values of left and right index
	array[left], array[right] = array[right], array[left]
	// Apply QuickSort to what's left of the array
	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return
}

// This implementation will sort the array and return the sorted one without modifying the original []int in parameter
func QuickSortGo(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSortGo(less), QuickSortGo(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}