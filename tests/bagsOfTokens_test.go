package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func Test(t *testing.T) {

	test := tester.Tester{T: t}

	tables := []struct {
		input1    	[]int
		input2 		int
		expected 	int
	}{
		{[]int{100, 200, 400, 300, 460, 760, 455, 324, 574}, 200, 3},
		{[]int{100, 200}, 150, 1},
		{[]int{48, 67, 26},81, 2},
		{[]int{100, 200, 400, 300, 460, 760, 455, 324, 574}, 50, 0},
	}

	for _, table := range tables {
		// Put the function you want to test here
		result := algorithms.BagOfTokensScore(table.input1, table.input2)
		if result != table.expected {
			test.Fail("Score of (%v) was INCORRECT, got: %d, should be: %d.",
				table.input1, result, table.expected)
		} else {
			test.Sucess("Score of (%v) was CORRECT, got : %d, should be: %d.",
				table.input1, result, table.expected)
		}
	}
}