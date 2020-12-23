package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func TestJumpGame3(t *testing.T) {

	test := tester.Tester{T: t}

	tables := []struct {
		arr    []int
		start  int
		expected bool
	}{
		{[]int{4,2,3,0,3,1,2}, 5, true},
		{[]int{4,2,3,0,3,1,2}, 0, true},
		{[]int{3,0,2,1,2}, 2, false},
	}

	for _, table := range tables {
		// Put the function you want to test here
		result := algorithms.CanReach(table.arr, table.start)
		if result != table.expected {
			test.Fail("(%v) Reachability of 0 was INCORRECT, got: %t, should be: %t.",
				table.arr, result, table.expected)
		} else {
			test.Sucess("(%v) Reachability of 0 was CORRECT, got : %t, should be: %t.",
				table.arr, result, table.expected)
		}
	}
}