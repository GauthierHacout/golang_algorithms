package algorithms

// Find the biggest difference possible between values of the array with addition/substitution of n (0<=n<=K)
func SmallestRange(A []int, K int) int {
	var diff int;
	min, max := minMax(A);

	diff = max - min - 2*K
	if diff<0 {
		return 0
	}
	return diff;
}

func minMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
