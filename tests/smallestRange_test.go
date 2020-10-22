package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func TestSmallestRange(t *testing.T) {

	test := tester.Tester{T: t}

	tables := []struct {
		input1    	[]int
		input2 		int
		expected 	int
	}{
		{[]int{7, 8, 8},5,0},
		{[]int{1,3,6, 18}, 4,9},
		{[]int{1,10},2,5},
		{[]int{1},0,0},
	}

	for _, table := range tables {
		// Put the function you want to test here
		result := algorithms.SmallestRange(table.input1, table.input2)
		if result != table.expected {
			test.Fail("Smallest range of (%v) was INCORRECT, got: %d, should be: %d.",
				table.input1, result, table.expected)
		} else {
			test.Sucess("Smallest range of (%v) was CORRECT, got : %d, should be: %d.",
				table.input1, result, table.expected)
		}
	}
}