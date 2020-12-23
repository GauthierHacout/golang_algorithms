package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func TestIncreasingArray( t *testing.T) {

	test := tester.Tester{T:t}

	tables := []struct {
			x []int
			y []int
			z int
		}{
			{[]int{1,5,3,6,7}, []int{1,3,2,4}, 1},
			{[]int{1,5,3,6,7}, []int{4,3,1}, 2},
			{[]int{1,5,3,6,7}, []int{1,6,3,3}, -1},
			{[]int{0,11,6,1,4,3}, []int{5,4,11,10,1,0}, 5},
			{[]int{5,16,19,2,1,12,7,14,5,16}, []int{6,17,4,3,6,13,4,3,18,17,16,17,7,14,1,16}, 8},
			{[]int{19,18,7,4,11,8,3,10,5,8,13}, []int{13,16,9,14,3,12,15,4,21,18,1,8,17,0,3,18}, 9},
		}

	for _, table := range tables {
		result := algorithms.MakeArrayIncreasing(table.x, table.y)
		if result != table.z {
			test.Fail("Increasing array of (%v) was INCORRECT, got: %d, should be: %d.", table.x, result, table.z)
		} else {
			test.Sucess("Increasing array of (%v) was CORRECT, got : %d, should be: %d.", table.x, result, table.z)
		}
	}
}
