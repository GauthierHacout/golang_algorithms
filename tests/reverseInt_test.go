package tests

import (
	"../algorithms"
	"../tester"
	"testing"
)

func TestReverseInt( t *testing.T) {

	test := tester.Tester{T:t}

	tables := []struct {
		x int
		y int
	}{
		{ 456, 654},
		{ -765, -567},
		{ 1534236469, 0},
		{ 10, 1},
	}

	for _, table := range tables {
		result := algorithms.ReverseInt(table.x)
		if result != table.y {
			test.Fail("Reversing of (%d) was INCORRECT, got: %d, should be: %d.", table.x, result, table.y)
		} else {
			test.Sucess("Reversing of (%d) was CORRECT, got : %d, should be: %d.", table.x, result, table.y)
		}
	}
}
