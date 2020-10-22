package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func TestKthSmallestInBST(t *testing.T) {

	test := tester.Tester{T: t}

	tables := []struct {
		input1    	*algorithms.TreeNode
		input2		int
		expected 	interface{}
	}{
		{algorithms.BSTfromSlice([]byte{3,1,4,2}), 1, 1},
		{algorithms.BSTfromSlice([]byte{5,3,6,2,4,1}), 3, 3},
		{algorithms.BSTfromSlice([]byte{5,3,6,2,4,1}), 67, "Not enough elements inside BST"},
	}

	for _, table := range tables {
		// Put the function you want to test here
		result, err := algorithms.KthSmallest(table.input1, table.input2)
		if err != nil {
			if err.Error() != table.expected {
				test.Fail("Kth smallest of (%v) was INCORRECT, got: %v, should be: %v.",
					table.input1, err.Error(), table.expected)
			} else {
				test.Sucess("Kth smallest of (%v) was CORRECT, got : %v, should be: %v.",
					table.input1, err.Error(), table.expected)
			}
		} else {
			if int(result) != table.expected {
				test.Fail("Kth smallest of (%v) was INCORRECT, got: %d, should be: %d.",
					table.input1, result, table.expected)
			} else {
				test.Sucess("Kth smallest of (%v) was CORRECT, got : %d, should be: %d.",
					table.input1, result, table.expected)
			}
		}
	}
}
