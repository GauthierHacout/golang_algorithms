package algorithms

import (
	"fmt"
	"math"
)

// Tried a recursive approach, it will always find a solution if it's possible but not necessarily the best one
// (with minimum amount of operations)
func MakeArrayIncreasing(arr1 []int, arr2 []int) int {
	//Remove duplicate from arr2
	arr2 = removeDuplicate(arr2)
	return recursiveSearchAndReplace(arr1, arr2, 0, 0);
}

func recursiveSearchAndReplace(arr1 []int, arr2 []int, i int, operations int) int {
	// Out condition
	if (i == len(arr1)-1){
		return operations
	}

	fmt.Printf(" ARRAY 1 : %v  \n", arr1)
	fmt.Printf(" ARRAY 2 : %v  \n", arr2)
	if (arr1[i] < arr1 [i+1]){
		//fmt.Printf("index: %d | value: %d | FINE \n", i, arr1[i]);
		//Go To Next Step
		//Make a copy of arr2 to be used in case of -1
		copyArr2 := make([]int, len(arr2))
		copy(copyArr2, arr2)
		copyArr1 := make([]int, len(arr1))
		copy(copyArr1, arr1)
		result := recursiveSearchAndReplace(arr1, arr2, i+1, operations);
		// OR Change arr1[i] to the smallest possible value in arr2
		// Why should I change arr1[i] because otherwise it's impossible so if recursiveSearchAndReplace returned -1
		// meaning it was impossible so I should change my approch and try again... -> meaning another loop of recursiveSearch
		// If it's impossible than return -1 anyway... It should work in theory but it's so underoptimized i'm crying
		// then I should try to modify arr[i] to the smallest value in arr2 if it's impossible -> -1
		if result == -1 {
			// arr2 is modified so it doesnt work create a copy of arr2 ?
			indexOfArr2, possible := searchFor(&copyArr2, -1)
			if possible {
				operations++
				copyArr1[i] = copyArr2[indexOfArr2]
				copyArr2[indexOfArr2] = -1
				return recursiveSearchAndReplace(copyArr1, copyArr2, i, operations);
			}
			return -1
		}
		return result
	}

	//fmt.Printf("index: %d | value: %d | --NOT-FINE-- \n", i, arr1[i])

	// Make this one smaller or make the next one bigger and compare each path than return the appropriate one
	var indexOfArr2 int
	var operationOccured bool
	if i==0 {
		indexOfArr2, operationOccured = searchFor(&arr2, -1)
	}else {
		indexOfArr2, operationOccured = searchFor(&arr2, arr1[i-1])
	}
	if operationOccured {
		//fmt.Printf("Operation Occured : %d \n", arr2[indexOfArr2])
		operations++
		if (arr1[i] <= arr2[indexOfArr2]) {
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
			//fmt.Printf(" Value found : %d | x : %d \n", (*arr2)[index], x)
		}
	}

	return index, index != -1
}

func removeDuplicate(slice []int) []int {
	sliceCopy := make([]int, 0)

	for _, value := range slice {
		if !sliceContainsValue(sliceCopy, value) {
			sliceCopy = append(sliceCopy, value)
		}
	}

	return sliceCopy
}

func sliceContainsValue(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
