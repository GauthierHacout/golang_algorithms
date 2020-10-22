package algorithms

import (
	"fmt"
	"math"
)

func MakeArrayIncreasing(arr1 []int, arr2 []int) int {

	return recursiveSearchAndReplace(arr1, arr2, 0, 0);
}


func recursiveSearchAndReplace(arr1 []int, arr2 []int, i int, operations int) int {
	// Out condition
	if (i == len(arr1)-1){
		return operations
	}

	fmt.Printf(" ARR1 : %v  \n", arr1)
	if (arr1[i] < arr1 [i+1]) || i==0 {
		fmt.Printf("index: %d | value: %d | FINE \n", i, arr1[i]);
		//Go To Next Step
		result := recursiveSearchAndReplace(arr1, arr2, i+1, operations);
		// OR Change arr1[i] to the smallest possible value in arr2
		// Why should I change arr1[i] because otherwise it's impossible so if recursiveSearchAndReplace returned -1
		// meaning it was impossible so I should change my approch and try again... -> meaning another loop of recursiveSearch
		// If it's impossible than return -1 anyway... It should work in theory but it's so underoptimized i'm crying
		// then I should try to modify arr[i] to the smallest value in arr2 if it's impossible -> -1
		if result == -1 {
			indexOfArr2, operationOccured := searchFor(&arr2, arr1[i])
			if operationOccured {
				operations++
				arr1[i] = arr2[indexOfArr2]
				return recursiveSearchAndReplace(arr1, arr2, i, operations);
			}
			return -1
		}
		return result
	}

	fmt.Printf("index: %d | value: %d | --NOT-FINE-- \n", i, arr1[i])
	indexOfArr2, operationOccured := searchFor(&arr2, arr1[i-1])
	if operationOccured {
		fmt.Printf("Operation Occured : %d \n", arr2[indexOfArr2])
		operations++
		if (arr1[i]<=arr2[indexOfArr2]) {
			arr1[i+1] = arr2[indexOfArr2]
		} else {
			arr1[i] = arr2[indexOfArr2]
		}
		arr2[indexOfArr2] = -1
		return recursiveSearchAndReplace(arr1, arr2, i, operations)
	}
	return -1
}


func searchFor(arr2 *[]int, x int) (int, bool) {

	var index int = -1
	var diff int = int(math.Pow10(9) + 1)

	for i := 0; i < len(*arr2); i++ {
		if ((*arr2)[i] > x && ((*arr2)[i]-x) < diff) {
			diff = (*arr2)[i] - x
			index = i
			fmt.Printf(" Value found : %d | x : %d \n", (*arr2)[index], x)
		}
	}

	return index, index != -1
}