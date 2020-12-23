package algorithms

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	result := head.Val

	for head.Next!=nil {
		result = (result << 1) | head.Next.Val
		head = head.Next
	}

	return result
}

func minCostToMoveChips(position []int) int {

	pair, impair := 0,0

	for i:=0; i<len(position); i++ {
		if position[i]%2==0 {
			pair++;
		} else {
			impair++;
		}
	}

	return int(math.Min(float64(pair), float64(impair)))
}
