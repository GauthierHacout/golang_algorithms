package tests

import (
	"../algorithms"
	"../tester"
	"reflect"
	"testing"
)

func TestQuickSort( t *testing.T) {

	test := tester.Tester{T: t}

	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{1,2,6,7,9,45,3},[]int{1,2,3,6,7,9,45}},
		{[]int{1,2,2,3,4,3},[]int{1,2,2,3,3,4}},
		{[]int{1,2,3,4,5},[]int{1,2,3,4,5}},
		{[]int{42,23},[]int{23,42}},
		{[]int{42},[]int{42}},
		{[]int{},[]int{}},
	}

	for _, table := range tables {
		result := algorithms.QuickSortGo(table.input)
		if !reflect.DeepEqual(result, table.expected) {
			test.Fail("Quicksort of (%v) was INCORRECT, got: %v, should be: %v.",
				table.input, result, table.expected)
		} else {
			test.Sucess("Quicksort of (%v) was CORRECT, got : %v, should be: %v.",
				table.input, result, table.expected)
		}
	}
}