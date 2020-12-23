package main

import "fmt"

func main() {

	/*test := []int{0,0,1,1,1,1,2,3}

	fmt.Println("SHOULD BE [0,0,1,1,2,3] : ",removeDuplicates(test))
	fmt.Println(test)*/

	test := []int{3,1,5,8}
	fmt.Println("SHOULD BE 167 : ", maxCoins(test))
	test1 := []int{9,76,64,21}
	fmt.Println("SHOULD BE 116718 : ", maxCoins(test1))
	test2 := []int{9,76,64,21,97,60}
	fmt.Println("SHOULD BE 1086136 : ", maxCoins(test2))

}
func maxCoins(nums []int) int {
	if nums==nil || len(nums)==0 {
		return 0
	}
	return calculateValue(nums, 0)
}

func calculateValue(nums []int, result int) int {
	if len(nums) == 1 {
		result += nums[0]
		return result
	}
	var value int
	var index int

	if len(nums)>4 {
		value, index = findBestPossibleValue(nums)
	} else {
		value, index = findMaxPossibleValue(nums)
	}
	result += value
	fmt.Println("\nNUMS : ", nums)
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("NUMS : ", nums)

	return calculateValue(nums, result)
}


func findMaxPossibleValue(nums []int) (int, int) {
	//Max value inside nums is 100
	var maxValue int
	var indexOfValue int

	for i:=0; i<len(nums); i++{
		var calculatedValue int
		if i==0 {
			calculatedValue = nums[i]*nums[i+1]
		} else if i==len(nums)-1 {
			calculatedValue = nums[i-1]*nums[i]
		} else {
			calculatedValue = nums[i-1]*nums[i]*nums[i+1]
		}

		if calculatedValue>maxValue {
			maxValue = calculatedValue
			indexOfValue = i
		}
	}



	return maxValue, indexOfValue
}


func findBestPossibleValue(nums []int) (int, int) {

	//Max value inside nums is 100
	minValue := 101
	var indexOfMinValue int

	for i:=1; i<len(nums)-1; i++{
		if nums[i]<minValue {
			minValue = nums[i]
			indexOfMinValue = i
		}
	}

	var calculatedValue int
	if indexOfMinValue==0 {
		calculatedValue = nums[indexOfMinValue]*nums[indexOfMinValue+1]
	} else if indexOfMinValue==len(nums)-1 {
		calculatedValue = nums[indexOfMinValue-1]*nums[indexOfMinValue]
	} else {
		calculatedValue = nums[indexOfMinValue-1]*nums[indexOfMinValue]*nums[indexOfMinValue+1]
	}

	return calculatedValue, indexOfMinValue
}

func removeDuplicates(nums []int) int {

	countOccurences := 0
	n := 0

	for i,x := range nums {

		if i != 0 {
			if x == nums[i-1] {
				countOccurences++
			} else {
				countOccurences = 0
			}
		}

		// We keep it
		if countOccurences<2 {
			nums[n] = x
			n++
		}
	}

	fmt.Println("NUMS : ", nums)
	fmt.Println(" N : ", n)
	nums = nums[:n]
	fmt.Println("NUMS : ", nums)
	return len(nums)
}